package v16

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

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
	FixHuahuaGauges(ctx, appKeepers.IncentivesKeeper, appKeepers.BankKeeper)
}

// Some gauges were incorrectly created for pools 605 (HUAHUA / OSMO) and 606 (HUAHUA / ATOM)
// and this error prevented them from distributing 6 billion HUAHUA tokens as incentives.
// This fix is here to release these incentives to their respective pools correctly.
func FixHuahuaGauges(ctx sdk.Context, incentiveskeeper *incentiveskeeper.Keeper, bank bankkeeper.Keeper) {
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
