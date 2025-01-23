package keepers

import (
	storetypes "cosmossdk.io/store/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	// ibc-go
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	// cosmwasm
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	betmoduletypes "github.com/sge-network/sge/x/bet/types"
	housemoduletypes "github.com/sge-network/sge/x/house/types"
	marketmoduletypes "github.com/sge-network/sge/x/market/types"
	minttypes "github.com/sge-network/sge/x/mint/types"
	orderbookmoduletypes "github.com/sge-network/sge/x/orderbook/types"
	ovmmoduletypes "github.com/sge-network/sge/x/ovm/types"
	rewardmoduletypes "github.com/sge-network/sge/x/reward/types"
	subaccounttypes "github.com/sge-network/sge/x/subaccount/types"
	// unnamed import of statik for swagger UI support
	// TODO: Check if needed
	// _ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

const (
	// AccountAddressPrefix prefix used for generating account address
	AccountAddressPrefix = "sge"
)

func (appKeepers *AppKeepers) GenerateKeys() {
	// Define what keys will be used in the cosmos-sdk key/value store.
	// Cosmos-SDK modules each have a "key" that allows the application to reference what they've stored on the chain.
	appKeepers.keys = storetypes.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey, crisistypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, consensusparamtypes.StoreKey, upgradetypes.StoreKey,
		feegrant.StoreKey, evidencetypes.StoreKey, capabilitytypes.StoreKey, authzkeeper.StoreKey, group.StoreKey,
		wasmtypes.StoreKey, ibcwasmtypes.StoreKey,
		ibcexported.StoreKey, ibctransfertypes.StoreKey, ibcfeetypes.StoreKey, icacontrollertypes.StoreKey, icahosttypes.StoreKey, ibchookstypes.StoreKey,
		orderbookmoduletypes.StoreKey,
		betmoduletypes.StoreKey,
		marketmoduletypes.StoreKey,
		ovmmoduletypes.StoreKey,
		housemoduletypes.StoreKey,
		rewardmoduletypes.StoreKey,
		subaccounttypes.StoreKey,
	)

	// Define transient store keys
	appKeepers.tkeys = storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)

	// MemKeys are for information that is stored only in RAM.
	appKeepers.memKeys = storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
}

func (appKeepers *AppKeepers) GetKVStoreKey() map[string]*storetypes.KVStoreKey {
	return appKeepers.keys
}

func (appKeepers *AppKeepers) GetTransientStoreKey() map[string]*storetypes.TransientStoreKey {
	return appKeepers.tkeys
}

func (appKeepers *AppKeepers) GetMemoryStoreKey() map[string]*storetypes.MemoryStoreKey {
	return appKeepers.memKeys
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetKey(storeKey string) *storetypes.KVStoreKey {
	return appKeepers.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return appKeepers.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (appKeepers *AppKeepers) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return appKeepers.memKeys[storeKey]
}
