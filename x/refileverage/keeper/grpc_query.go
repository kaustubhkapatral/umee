package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/umee-network/umee/v5/x/refileverage/types"
)

var _ types.QueryServer = Querier{}

// Querier implements a QueryServer for the x/refileverage module.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) Params(
	goCtx context.Context,
	req *types.QueryParams,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.Keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q Querier) RegisteredTokens(
	goCtx context.Context,
	req *types.QueryRegisteredTokens,
) (*types.QueryRegisteredTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var tokens []types.Token
	if len(req.BaseDenom) != 0 {
		token, err := q.Keeper.GetTokenSettings(ctx, req.BaseDenom)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	} else {
		tokens = q.Keeper.GetAllRegisteredTokens(ctx)
	}

	return &types.QueryRegisteredTokensResponse{
		Registry: tokens,
	}, nil
}

func (q Querier) MarketSummary(
	goCtx context.Context,
	req *types.QueryMarketSummary,
) (*types.QueryMarketSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "empty denom")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	token, err := q.Keeper.GetTokenSettings(ctx, req.Denom)
	if err != nil {
		return nil, err
	}

	rate := q.Keeper.DeriveExchangeRate(ctx, req.Denom)
	borrowAPY := q.Keeper.DeriveBorrowAPY(ctx, req.Denom)

	supplied, _ := q.Keeper.GetTotalSupply(ctx, req.Denom)
	balance := q.Keeper.ModuleBalance(ctx, req.Denom).Amount
	reserved := q.Keeper.GetReserves(ctx, req.Denom).Amount
	borrowed := q.Keeper.GetTotalBorrowed(ctx)

	uDenom := types.ToUTokenDenom(req.Denom)
	uSupply := q.Keeper.GetUTokenSupply(ctx, uDenom)
	uCollateral := q.Keeper.GetTotalCollateral(ctx, uDenom)

	// availableCollateralize respects both MaxCollateralShare and MinCollateralLiquidity
	maxCollateral, _ := q.Keeper.maxCollateralFromShare(ctx, uDenom)
	availableCollateralize := maxCollateral.Sub(uCollateral.Amount)
	availableCollateralize = sdk.MaxInt(availableCollateralize, sdk.ZeroInt())

	resp := types.QueryMarketSummaryResponse{
		SymbolDenom:            token.SymbolDenom,
		Exponent:               token.Exponent,
		UTokenExchangeRate:     rate,
		Borrow_APY:             borrowAPY,
		Supplied:               supplied.Amount,
		Reserved:               reserved,
		Collateral:             uCollateral.Amount,
		Borrowed:               borrowed,
		Liquidity:              balance.Sub(reserved),
		MaximumCollateral:      maxCollateral,
		UTokenSupply:           uSupply.Amount,
		AvailableCollateralize: availableCollateralize,
	}

	// Oracle prices in response will be nil if it is unavailable
	oraclePrice, _, oracleErr := q.Keeper.TokenPrice(ctx, req.Denom, types.PriceModeSpot)
	if oracleErr == nil {
		resp.OraclePrice = &oraclePrice
	} else {
		resp.Errors += oracleErr.Error()
	}
	historicPrice, _, historicErr := q.Keeper.TokenPrice(ctx, req.Denom, types.PriceModeHistoric)
	if historicErr == nil {
		resp.OracleHistoricPrice = &historicPrice
	} else {
		resp.Errors += historicErr.Error()
	}

	return &resp, nil
}

func (q Querier) AccountBalances(
	goCtx context.Context,
	req *types.QueryAccountBalances,
) (*types.QueryAccountBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	supplied, err := q.Keeper.GetAllSupplied(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &types.QueryAccountBalancesResponse{
		Supplied:   supplied,
		Collateral: q.Keeper.GetBorrowerCollateral(ctx, addr),
		Borrowed:   q.Keeper.GetBorrowerBorrows(ctx, addr),
	}, nil
}

func (q Querier) AccountSummary(
	goCtx context.Context,
	req *types.QueryAccountSummary,
) (*types.QueryAccountSummaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	// supplied, err := q.Keeper.GetAllSupplied(ctx, addr)
	// if err != nil {
	// 	return nil, err
	// }
	collateral := q.Keeper.GetBorrowerCollateral(ctx, addr)
	borrowed := q.Keeper.GetBorrowerBorrows(ctx, addr)

	// collateral value always uses spot prices, and this line skips assets that are missing prices
	collateralValue, err := q.Keeper.VisibleCollateralValue(ctx, collateral)
	if err != nil {
		return nil, err
	}

	// borrow limit shown here as it is used in refileverage logic:
	// using the lower of spot or historic prices for each collateral token
	// skips collateral tokens with missing oracle prices
	borrowLimit, err := q.Keeper.VisibleBorrowLimit(ctx, collateral)
	if err != nil {
		return nil, err
	}

	resp := &types.QueryAccountSummaryResponse{
		CollateralValue: collateralValue,
		BorrowedValue:   borrowed,
		BorrowLimit:     borrowLimit,
	}

	// liquidation always uses spot prices. This response field will be null
	// if a price is missing
	liquidationThreshold, err := q.Keeper.CalculateLiquidationThreshold(ctx, collateral)
	if err == nil {
		resp.LiquidationThreshold = &liquidationThreshold
	}

	return resp, nil
}

func (q Querier) LiquidationTargets(
	goCtx context.Context,
	req *types.QueryLiquidationTargets,
) (*types.QueryLiquidationTargetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if !q.Keeper.liquidatorQueryEnabled {
		return nil, types.ErrNotLiquidatorNode
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	targets, err := q.Keeper.GetEligibleLiquidationTargets(ctx)
	if err != nil {
		return nil, err
	}

	stringTargets := []string{}
	for _, addr := range targets {
		stringTargets = append(stringTargets, addr.String())
	}

	return &types.QueryLiquidationTargetsResponse{Targets: stringTargets}, nil
}

func (q Querier) BadDebts(
	goCtx context.Context,
	req *types.QueryBadDebts,
) (*types.QueryBadDebtsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	targets := q.Keeper.getAllBadDebts(ctx)

	return &types.QueryBadDebtsResponse{Targets: targets}, nil
}

func (q Querier) MaxWithdraw(
	goCtx context.Context,
	req *types.QueryMaxWithdraw,
) (*types.QueryMaxWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	denoms := []string{}
	maxUTokens := sdk.NewCoins()
	maxTokens := sdk.NewCoins()

	if req.Denom != "" {
		// Denom specified
		denoms = []string{req.Denom}
	} else {
		// Denom not specified
		for _, t := range q.Keeper.GetAllRegisteredTokens(ctx) {
			if !t.Blacklist {
				denoms = append(denoms, t.BaseDenom)
			}
		}
	}

	for _, denom := range denoms {
		// If a price is missing for the borrower's collateral,
		// but not this uToken or any of their borrows, error
		// will be nil and the resulting value will be what
		// can safely be withdrawn even with missing prices.
		// On non-nil error here, max withdraw is zero.
		uToken, _, err := q.Keeper.userMaxWithdraw(ctx, addr, denom)
		if err == nil && uToken.IsPositive() {
			token, err := q.Keeper.ExchangeUToken(ctx, uToken)
			if err != nil {
				return nil, err
			}
			maxUTokens = maxUTokens.Add(uToken)
			maxTokens = maxTokens.Add(token)
		}
		// Non-price errors will cause the query itself to fail
		if nonOracleError(err) {
			return nil, err
		}
	}

	return &types.QueryMaxWithdrawResponse{
		Tokens:  maxTokens,
		UTokens: maxUTokens,
	}, nil
}

func (q Querier) MaxBorrow(
	goCtx context.Context,
	req *types.QueryMaxBorrow,
) (*types.QueryMaxBorrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "empty address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	denoms := []string{}
	maxTokens := sdk.NewCoins()

	if req.Denom != "" {
		// Denom specified
		denoms = []string{req.Denom}
	} else {
		// Denom not specified
		for _, t := range q.Keeper.GetAllRegisteredTokens(ctx) {
			if !t.Blacklist {
				denoms = append(denoms, t.BaseDenom)
			}
		}
	}

	for _, denom := range denoms {
		// If a price is missing for the borrower's collateral,
		// but not this token or any of their borrows, error
		// will be nil and the resulting value will be what
		// can safely be borrowed even with missing prices.
		// On non-nil error here, max borrow is zero.
		maxBorrow, err := q.Keeper.userMaxBorrow(ctx, addr, denom)
		if err == nil && maxBorrow.IsPositive() {
			maxTokens = maxTokens.Add(maxBorrow)
		}
		// Non-price errors will cause the query itself to fail
		if nonOracleError(err) {
			return nil, err
		}
	}

	return &types.QueryMaxBorrowResponse{
		Tokens: maxTokens,
	}, nil
}