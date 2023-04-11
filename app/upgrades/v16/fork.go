package v16

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	incentiveskeeper "github.com/osmosis-labs/osmosis/v15/x/incentives/keeper"

	"github.com/osmosis-labs/osmosis/v15/app/keepers"
)

// RunForkLogic executes height-gated on-chain fork logic for the Osmosis v16
// upgrade.
func RunForkLogic(ctx sdk.Context, appKeepers *keepers.AppKeepers) {
	for i := 0; i < 100; i++ {
		ctx.Logger().Info("I am upgrading to v16")
	}
	ctx.Logger().Info("Applying Osmosis v16 upgrade.")
	ctx.Logger().Info("Fixing the Huahua gauges.")
	FixHuahuaGauges(ctx, appKeepers.IncentivesKeeper)
}

// Some gauges were incorrectly created for pools 605 (HUAHUA / OSMO) and 606 (HUAHUA / ATOM)
// and this error prevented them from distributing 6 billion HUAHUA tokens as incentives.
// This fix is here to release these incentives to their respective pools correctly.
func FixHuahuaGauges(ctx sdk.Context, incentiveskeeper *incentiveskeeper.Keeper) {
	// Huahua incorrect gauges are IDs 1954 to 1959
	gaugeIDs := []uint64{1954, 1955, 1956, 1957, 1958, 1959}

	gauges, err := incentiveskeeper.GetGaugeFromIDs(ctx, gaugeIDs)

	if err != nil {
		ctx.Logger().Error("v16 upgrade: Could not find the Huahua gauges for fixing")
		return
	}

	// Fix the QueryCondition Denom
	for _, gauge := range gauges {
		if gauge.DistributeTo.Denom == "GAMM605" {
			gauge.DistributeTo.Denom = "pool/gamm/605"
		} else if gauge.DistributeTo.Denom == "GAMM606" {
			gauge.DistributeTo.Denom = "pool/gamm/606"
		}
	}

	// Overwrite the gauges in the store
	incentiveskeeper.OverwriteGaugeV16MigrationUnsafe(ctx, gauges)
}
