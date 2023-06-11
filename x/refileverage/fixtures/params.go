package fixtures

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v5/x/refileverage/types"
)

// Params returns refileverage params used in testing.
func Params() types.Params {
	return types.Params{
		CompleteLiquidationThreshold: sdk.MustNewDecFromStr("0.1"),
		MinimumCloseFactor:           sdk.MustNewDecFromStr("0.01"),
		OracleRewardFactor:           sdk.MustNewDecFromStr("0.01"),
		SmallLiquidationSize:         sdk.MustNewDecFromStr("100.00"),
		DirectLiquidationFee:         sdk.MustNewDecFromStr("0.1"),
	}
}