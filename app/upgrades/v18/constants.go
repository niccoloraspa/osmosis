package v18

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/osmosis-labs/osmosis/v17/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v18 upgrade.
const UpgradeName = "v18"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}