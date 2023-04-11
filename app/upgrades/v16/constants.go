package v16

import (
	"github.com/osmosis-labs/osmosis/v15/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"

	cltypes "github.com/osmosis-labs/osmosis/v15/x/concentrated-liquidity/types"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v16 upgrade.
const UpgradeName = "v16"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{cltypes.StoreKey},
		Deleted: []string{},
	},
}

// ForkHeight defines the block height at which the Osmosis v16 upgrade is
// triggered.
// Block height [Write block height as X_XXX_XXX when decided], approximately [Write approximate UTC time here when decided]
const ForkHeight = 9999999 // TODO: set the appropriate height when decided

var Fork = upgrades.Fork{
	UpgradeName:    UpgradeName,
	UpgradeHeight:  ForkHeight,
	BeginForkLogic: RunForkLogic,
}
