package v16

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/osmosis-labs/osmosis/v15/app/keepers"
	"github.com/osmosis-labs/osmosis/v15/app/upgrades"

	incentiveskeeper "github.com/osmosis-labs/osmosis/v15/x/incentives/keeper"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		ctx.Logger().Info("Applying Osmosis v16 upgrade.")
		ctx.Logger().Info("Fixing the Huahua gauges.")
		fixHuahuaGauges(ctx, keepers.IncentivesKeeper, keepers.BankKeeper)

		return migrations, nil
	}
}

// Some gauges were incorrectly created for pools 605 (HUAHUA / OSMO) and 606 (HUAHUA / ATOM)
// and this error prevented them from distributing 6 billion HUAHUA tokens as incentives.
// This fix is here to release the coins locked in gauges to the Huahua team wallet and archive these gauges
func fixHuahuaGauges(ctx sdk.Context, incentiveskeeper *incentiveskeeper.Keeper, bank bankkeeper.Keeper) {
	// Huahua incorrect gauges are IDs 1954 to 1959
	gaugeIDs := []uint64{1954, 1955, 1956, 1957, 1958, 1959}

	// Huahua Team wallet is osmo14fketv99hlrlk80mkggw643spsj3yyf7t2pjhr
	const huahuaTeamWalletAddrString = "osmo14fketv99hlrlk80mkggw643spsj3yyf7t2pjhr"

	huahuaTeamWalletAddr, err := sdk.AccAddressFromBech32(huahuaTeamWalletAddrString)
	if err != nil {
		ctx.Logger().Error("v16 upgrade: Could not parse Huahua team wallet")
		panic(err)
	}

	gauges, err := incentiveskeeper.GetGaugeFromIDs(ctx, gaugeIDs)

	if err != nil {
		ctx.Logger().Error("v16 upgrade: Could not find the Huahua gauges for fixing")
		panic(err)
	}

	allCoins := sdk.Coins{}

	for _, gauge := range gauges {
		// make sure that we selected the Huahua gauges that have incorrect QueryCondition (DistributeTo.Denom)
		if gauge.DistributeTo.Denom != "GAMM605" && gauge.DistributeTo.Denom != "GAMM606" {
			panic("v16 upgrade: Didn't retrieve expected Huahua gauges")
		}

		// retrieve the exact amount of coins from the gauge
		allCoins = allCoins.Add(gauge.Coins...)
		// reset the gauge to no coin
		gauge.Coins = sdk.Coins{}
		// mark the gauge as filled over the required num of epochs
		gauge.FilledEpochs = gauge.NumEpochsPaidOver
	}

	// Overwrite the gauges in the store
	err = incentiveskeeper.OverwriteGaugeV16MigrationUnsafe(ctx, gauges)
	if err != nil {
		ctx.Logger().Error("v16 upgrade: Could not overwrite Huahua gauges")
		panic(err)
	}

	err = bank.SendCoinsFromModuleToAccount(ctx, "incentives", huahuaTeamWalletAddr, allCoins)
	if err != nil {
		ctx.Logger().Error("v16 upgrade: Could not send coins to Huahua team wallet")
		panic(err)
	}
}
