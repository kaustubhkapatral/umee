package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v6/x/oracle/types"
)

// GetOracleAccount returns oracle ModuleAccount.
func (k Keeper) GetOracleAccount(ctx sdk.Context) sdk.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

// GetRewardPool retrieves the balance of the oracle module account.
func (k Keeper) GetRewardPool(ctx sdk.Context, denom string) sdk.Coin {
	acc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	return k.bankKeeper.GetBalance(ctx, acc.GetAddress(), denom)
}
