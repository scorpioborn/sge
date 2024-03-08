package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sge-network/sge/utils"
	"github.com/sge-network/sge/x/market/types"
)

// Keeper is the type for module properties
type Keeper struct {
	cdc             codec.BinaryCodec
	storeKey        storetypes.StoreKey
	memKey          storetypes.StoreKey
	paramStore      paramtypes.Subspace
	modFunder       *utils.ModuleAccFunder
	ovmKeeper       types.OVMKeeper
	orderbookKeeper types.OrderbookKeeper
}

// SdkExpectedKeepers contains expected keepers parameter needed by NewKeeper
type SdkExpectedKeepers struct {
	BankKeeper    types.BankKeeper
	AccountKeeper types.AccountKeeper
	AuthzKeeper   types.AuthzKeeper
}

// NewKeeper creates new keeper object
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	expectedKeepers SdkExpectedKeepers,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramStore: ps,
		modFunder: utils.NewModuleAccFunder(
			expectedKeepers.BankKeeper,
			expectedKeepers.AccountKeeper,
			types.ErrorBank,
		),
	}
}

// SetOrderbookKeeper sets the orderbook module keeper to the market keeper.
func (k *Keeper) SetOrderbookKeeper(orderBookKeeper types.OrderbookKeeper) {
	k.orderbookKeeper = orderBookKeeper
}

// SetOVMKeeper sets the ovm module keeper to the market keeper.
func (k *Keeper) SetOVMKeeper(ovmKeeper types.OVMKeeper) {
	k.ovmKeeper = ovmKeeper
}

// Logger returns the logger of the keeper
func (Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
