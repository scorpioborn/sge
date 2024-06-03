package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sge-network/sge/testutil/simapp"
	"github.com/sge-network/sge/x/subaccount/keeper"
	"github.com/sge-network/sge/x/subaccount/types"
)

func setupMsgServerAndApp(
	t testing.TB,
) (*simapp.TestApp, *keeper.Keeper, types.MsgServer, sdk.Context) {
	tApp, k, ctx := setupKeeperAndApp(t)
	return tApp, k, keeper.NewMsgServerImpl(*k), ctx
}
