package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sge-network/sge/consts"
	"github.com/sge-network/sge/x/mint/types"
)

// Keeper is the type for module properties
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	paramstore       paramtypes.Subspace
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	feeCollectorName string
	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// ExpectedKeepers contains expected keepers parameter needed by NewKeeper
type ExpectedKeepers struct {
	StakingKeeper types.StakingKeeper
	BankKeeper    types.BankKeeper
}

// NewKeeper creates new keeper object
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	expectedKeepers ExpectedKeepers,
	feeCollectorName string,
	authority string,
) *Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf(consts.ErrModuleAccountHasNotBeenSet, types.ModuleName))
	}

	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		paramstore:       ps,
		stakingKeeper:    expectedKeepers.StakingKeeper,
		bankKeeper:       expectedKeepers.BankKeeper,
		feeCollectorName: feeCollectorName,
		authority:        authority,
	}
}

// Logger returns the logger of the keeper
func (Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
