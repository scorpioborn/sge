package v10

import (
	"context"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sge-network/sge/app/keepers"
	minttypes "github.com/sge-network/sge/x/mint/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	k *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		amount := sdk.NewCoins(sdk.NewCoin("uusdc", math.NewInt(int64(1000000000000000000))))
		amountSge := sdk.NewCoins(sdk.NewCoin("usge", math.NewInt(int64(1000000000000))))
		receivingAddress := sdk.MustAccAddressFromBech32("sge1989lnphmjyts54d84tdvl799dy9k3604f45h2l")
		k.MintKeeper.MintCoins(ctx, amount)
		k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, receivingAddress, amount)

		k.MintKeeper.MintCoins(ctx, amountSge)
		k.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, receivingAddress, amountSge)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
