package keepers

import (
	storetypes "cosmossdk.io/store/types"
	circuitkeeper "cosmossdk.io/x/circuit/keeper"
	circuittypes "cosmossdk.io/x/circuit/types"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantkeeper "cosmossdk.io/x/feegrant/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	sdkparams "github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icacontroller "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"

	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v8/keeper"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	ibc_hooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/keeper"
	ibcwasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/keeper"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
	icahost "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	packetforwardkeeper "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/keeper"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"

	// IBC Transfer: Defines the "transfer" IBC port
	transfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"

	// cosmwasm
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	// sge

	"github.com/sge-network/sge/app/params"
	mintkeeper "github.com/sge-network/sge/x/mint/keeper"
	minttypes "github.com/sge-network/sge/x/mint/types"

	betmodule "github.com/sge-network/sge/x/bet"
	betmodulekeeper "github.com/sge-network/sge/x/bet/keeper"
	betmoduletypes "github.com/sge-network/sge/x/bet/types"

	marketmodule "github.com/sge-network/sge/x/market"
	marketmodulekeeper "github.com/sge-network/sge/x/market/keeper"
	marketmoduletypes "github.com/sge-network/sge/x/market/types"

	ovmmodule "github.com/sge-network/sge/x/ovm"
	ovmmodulekeeper "github.com/sge-network/sge/x/ovm/keeper"
	ovmmoduletypes "github.com/sge-network/sge/x/ovm/types"

	housemodule "github.com/sge-network/sge/x/house"
	housemodulekeeper "github.com/sge-network/sge/x/house/keeper"
	housemoduletypes "github.com/sge-network/sge/x/house/types"

	orderbookmodule "github.com/sge-network/sge/x/orderbook"
	orderbookmodulekeeper "github.com/sge-network/sge/x/orderbook/keeper"
	orderbookmoduletypes "github.com/sge-network/sge/x/orderbook/types"

	subaccountmodule "github.com/sge-network/sge/x/subaccount"
	subaccountmodulekeeper "github.com/sge-network/sge/x/subaccount/keeper"
	subaccountmoduletypes "github.com/sge-network/sge/x/subaccount/types"

	rewardmodule "github.com/sge-network/sge/x/reward"
	rewardmodulekeeper "github.com/sge-network/sge/x/reward/keeper"
	rewardmoduletypes "github.com/sge-network/sge/x/reward/types"
	// unnamed import of statik for swagger UI support
	// TODO: Check if needed
	// _ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	ParamsKeeper          *paramskeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ConsensusParamsKeeper *consensusparamkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper
	ScopedICQKeeper           capabilitykeeper.ScopedKeeper

	// "Normal" keepers
	AccountKeeper        *authkeeper.AccountKeeper
	BankKeeper           *bankkeeper.BaseKeeper
	AuthzKeeper          *authzkeeper.Keeper
	FeeGrantKeeper       *feegrantkeeper.Keeper
	GroupKeeper          *groupkeeper.Keeper
	StakingKeeper        *stakingkeeper.Keeper
	DistrKeeper          *distrkeeper.Keeper
	SlashingKeeper       *slashingkeeper.Keeper
	IBCKeeper            *ibckeeper.Keeper
	ICAHostKeeper        *icahostkeeper.Keeper
	ICAControllerKeeper  *icacontrollerkeeper.Keeper
	ICQKeeper            *icqkeeper.Keeper
	TransferKeeper       *ibctransferkeeper.Keeper
	IBCWasmClientKeeper  *ibcwasmkeeper.Keeper
	EvidenceKeeper       *evidencekeeper.Keeper
	GovKeeper            *govkeeper.Keeper
	WasmKeeper           *wasmkeeper.Keeper
	ContractKeeper       *wasmkeeper.PermissionedKeeper
	CircuitBreakerKeeper *circuitkeeper.Keeper
	IBCFeeKeeper         ibcfeekeeper.Keeper
	IBCHooksKeeper       *ibchookskeeper.Keeper

	// IBC modules
	// transfer module
	RawIcs20TransferAppModule transfer.AppModule
	TransferStack             *ibc_hooks.IBCMiddleware
	Ics20WasmHooks            *ibc_hooks.WasmHooks
	HooksICS4Wrapper          ibc_hooks.ICS4Middleware
	PacketForwardKeeper       *packetforwardkeeper.Keeper

	BetKeeper        *betmodulekeeper.Keeper
	MarketKeeper     *marketmodulekeeper.Keeper
	MintKeeper       *mintkeeper.Keeper
	HouseKeeper      *housemodulekeeper.Keeper
	OrderbookKeeper  *orderbookmodulekeeper.Keeper
	OVMKeeper        *ovmmodulekeeper.Keeper
	RewardKeeper     *rewardmodulekeeper.Keeper
	SubaccountKeeper *subaccountmodulekeeper.Keeper

	BetModule        *betmodule.AppModule
	MarketModule     *marketmodule.AppModule
	HouseModule      *housemodule.AppModule
	OrderbookModule  *orderbookmodule.AppModule
	OVMModule        *ovmmodule.AppModule
	RewardModule     *rewardmodule.AppModule
	SubaccountModule *subaccountmodule.AppModule
}

// InitSpecialKeepers initiates special keepers (crisis appkeeper, upgradekeeper, params keeper)
func (appKeepers *AppKeepers) InitSpecialKeepers(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	wasmDir string,
	cdc *codec.LegacyAmino,
	invCheckPeriod uint,
	skipUpgradeHeights map[int64]bool,
	homePath string,
) {
	appKeepers.GenerateKeys()
	paramsKeeper := appKeepers.initParamsKeeper(appCodec, cdc, appKeepers.keys[paramstypes.StoreKey], appKeepers.tkeys[paramstypes.TStoreKey])
	appKeepers.ParamsKeeper = &paramsKeeper

	// set the BaseApp's parameter store
	consensusParamsKeeper := consensusparamkeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(appKeepers.keys[consensusparamtypes.StoreKey]), authtypes.NewModuleAddress(govtypes.ModuleName).String(), runtime.EventService{})
	appKeepers.ConsensusParamsKeeper = &consensusParamsKeeper
	bApp.SetParamStore(appKeepers.ConsensusParamsKeeper.ParamsStore)

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, appKeepers.keys[capabilitytypes.StoreKey], appKeepers.memKeys[capabilitytypes.MemStoreKey])
	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	appKeepers.ScopedICAHostKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	appKeepers.ScopedICAControllerKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	appKeepers.ScopedWasmKeeper = appKeepers.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)
	appKeepers.ScopedICQKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icqtypes.ModuleName)
	appKeepers.CapabilityKeeper.Seal()

	// TODO: Make a SetInvCheckPeriod fn on CrisisKeeper.
	// IMO, its bad design atm that it requires this in state machine initialization
	crisisKeeper := crisiskeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(appKeepers.keys[crisistypes.StoreKey]), invCheckPeriod, appKeepers.BankKeeper, authtypes.FeeCollectorName, authtypes.NewModuleAddress(govtypes.ModuleName).String(), addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()))
	appKeepers.CrisisKeeper = crisisKeeper

	upgradeKeeper := upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(appKeepers.keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		bApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.UpgradeKeeper = upgradeKeeper
}

// InitNormalKeepers initializes all 'normal' keepers (account, app, bank, auth, staking, distribution, slashing, transfer, gamm, IBC router, pool incentives, governance, mint, txfees keepers).
func (appKeepers *AppKeepers) InitNormalKeepers(
	appCodec codec.Codec,
	encodingConfig params.EncodingConfig,
	bApp *baseapp.BaseApp,
	maccPerms map[string][]string,
	dataDir string,
	wasmDir string,
	wasmConfig wasmtypes.WasmConfig,
	wasmOpts []wasmkeeper.Option,
	blockedAddress map[string]bool,
	ibcWasmConfig ibcwasmtypes.WasmConfig,
) {
	govModAddress := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	legacyAmino := encodingConfig.Amino
	// Add 'normal' keepers
	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		AccountAddressPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.AccountKeeper = &accountKeeper
	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[banktypes.StoreKey]),
		appKeepers.AccountKeeper,
		blockedAddress,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		bApp.Logger(),
	)
	appKeepers.BankKeeper = &bankKeeper

	authzKeeper := authzkeeper.NewKeeper(
		runtime.NewKVStoreService(appKeepers.keys[authzkeeper.StoreKey]),
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper,
	)
	appKeepers.AuthzKeeper = &authzKeeper

	groupConfig := group.DefaultConfig()
	groupKeeper := groupkeeper.NewKeeper(
		appKeepers.keys[group.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper, groupConfig,
	)
	appKeepers.GroupKeeper = &groupKeeper

	feeGrantKeeper := feegrantkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[feegrant.StoreKey]),
		appKeepers.AccountKeeper,
	)
	appKeepers.FeeGrantKeeper = &feeGrantKeeper

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[stakingtypes.StoreKey]),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)
	appKeepers.StakingKeeper = stakingKeeper

	distrKeeper := distrkeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(appKeepers.keys[distrtypes.StoreKey]),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.DistrKeeper = &distrKeeper

	slashingKeeper := slashingkeeper.NewKeeper(
		appCodec,
		legacyAmino,
		runtime.NewKVStoreService(appKeepers.keys[slashingtypes.StoreKey]),
		appKeepers.StakingKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.SlashingKeeper = &slashingKeeper

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibchost.StoreKey],
		appKeepers.GetSubspace(ibchost.ModuleName),
		appKeepers.StakingKeeper,
		appKeepers.UpgradeKeeper,
		appKeepers.ScopedIBCKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// // Configure the hooks keeper
	// hooksKeeper := ibchookskeeper.NewKeeper(
	// 	appKeepers.keys[ibchookstypes.StoreKey],
	// 	appKeepers.GetSubspace(ibchookstypes.ModuleName),
	// 	appKeepers.IBCKeeper.ChannelKeeper,
	// 	nil,
	// )
	// appKeepers.IBCHooksKeeper = hooksKeeper

	// We are using a separate VM here
	ibcWasmClientKeeper := ibcwasmkeeper.NewKeeperWithConfig(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[ibcwasmtypes.StoreKey]),
		appKeepers.IBCKeeper.ClientKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		ibcWasmConfig,
		bApp.GRPCQueryRouter(),
	)

	appKeepers.IBCWasmClientKeeper = &ibcWasmClientKeeper

	// appKeepers.WireICS20PreWasmKeeper(appCodec, bApp, appKeepers.IBCHooksKeeper)

	icaHostKeeper := icahostkeeper.NewKeeper(
		appCodec, appKeepers.keys[icahosttypes.StoreKey],
		appKeepers.GetSubspace(icahosttypes.SubModuleName),
		appKeepers.IBCFeeKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.ScopedICAHostKeeper,
		bApp.MsgServiceRouter(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	icaHostKeeper.WithQueryRouter(bApp.GRPCQueryRouter())
	appKeepers.ICAHostKeeper = &icaHostKeeper

	icaControllerKeeper := icacontrollerkeeper.NewKeeper(
		appCodec, appKeepers.keys[icacontrollertypes.StoreKey],
		appKeepers.GetSubspace(icacontrollertypes.SubModuleName),
		appKeepers.IBCFeeKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedICAControllerKeeper,
		bApp.MsgServiceRouter(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.ICAControllerKeeper = &icaControllerKeeper

	// initialize ICA module with mock module as the authentication module on the controller side
	var icaControllerStack porttypes.IBCModule
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, *appKeepers.ICAControllerKeeper)

	// RecvPacket, message that originates from core IBC and goes down to app, the flow is:
	// channel.RecvPacket -> fee.OnRecvPacket -> icaHost.OnRecvPacket
	icaHostStack := icahost.NewIBCModule(*appKeepers.ICAHostKeeper)

	// ICQ Keeper
	icqKeeper := icqkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[icqtypes.StoreKey],
		appKeepers.IBCKeeper.ChannelKeeper, // may be replaced with middleware
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedICQKeeper,
		bApp.GRPCQueryRouter(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	appKeepers.ICQKeeper = &icqKeeper

	// Create Async ICQ module
	icqModule := icq.NewIBCModule(*appKeepers.ICQKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		// The transferIBC module is replaced by rateLimitingTransferModule
		AddRoute(ibctransfertypes.ModuleName, appKeepers.TransferStack).
		// Add icq modules to IBC router
		AddRoute(icqtypes.ModuleName, icqModule)
	// Note: the sealing is done after creating wasmd and wiring that up

	// create evidence keeper with router
	// If evidence needs to be handled for the app, set routes in router here and seal
	appKeepers.EvidenceKeeper = evidencekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[evidencetypes.StoreKey]),
		appKeepers.StakingKeeper,
		appKeepers.SlashingKeeper,
		addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr),
		runtime.ProvideCometInfoService(),
	)

	circuitBreakerKeeper := circuitkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[circuittypes.StoreKey]),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
	)
	appKeepers.CircuitBreakerKeeper = &circuitBreakerKeeper

	mintKeeper := mintkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[minttypes.StoreKey],
		appKeepers.GetSubspace(minttypes.ModuleName),
		appKeepers.AccountKeeper,
		mintkeeper.ExpectedKeepers{
			StakingKeeper: appKeepers.StakingKeeper,
			BankKeeper:    appKeepers.BankKeeper,
		},
		authtypes.FeeCollectorName,
	)
	appKeepers.MintKeeper = mintKeeper

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	wasmCapabilities := wasmkeeper.BuiltInCapabilities()

	// wasmOpts = append(owasm.RegisterCustomPlugins(appKeepers.BankKeeper, appKeepers.TokenFactoryKeeper), wasmOpts...)
	// wasmOpts = append(owasm.RegisterStargateQueries(*bApp.GRPCQueryRouter(), appCodec), wasmOpts...)

	wasmKeeper := wasmkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(appKeepers.keys[wasmtypes.StoreKey]),
		*appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		*appKeepers.StakingKeeper,
		distrkeeper.NewQuerier(*appKeepers.DistrKeeper),
		appKeepers.IBCFeeKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedWasmKeeper,
		appKeepers.TransferKeeper,
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		wasmCapabilities,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		wasmOpts...,
	)
	appKeepers.WasmKeeper = &wasmKeeper

	// Pass the contract keeper to all the structs (generally ICS4Wrappers for ibc middlewares) that need it
	// set the contract keeper for the Ics20WasmHooks
	appKeepers.ContractKeeper = wasmkeeper.NewDefaultPermissionKeeper(appKeepers.WasmKeeper)
	appKeepers.Ics20WasmHooks.ContractKeeper = appKeepers.WasmKeeper

	appKeepers.OrderbookKeeper = orderbookmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[orderbookmoduletypes.StoreKey],
		appKeepers.GetSubspace(orderbookmoduletypes.ModuleName),
		orderbookmodulekeeper.SdkExpectedKeepers{
			BankKeeper:     appKeepers.BankKeeper,
			AccountKeeper:  appKeepers.AccountKeeper,
			FeeGrantKeeper: appKeepers.FeeGrantKeeper,
		},
		govModAddress,
	)

	appKeepers.OVMKeeper = ovmmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ovmmoduletypes.StoreKey],
		appKeepers.keys[ovmmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(ovmmoduletypes.ModuleName),
		govModAddress,
	)

	appKeepers.MarketKeeper = marketmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[marketmoduletypes.StoreKey],
		appKeepers.keys[marketmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(marketmoduletypes.ModuleName),
		govModAddress,
	)
	appKeepers.MarketKeeper.SetOVMKeeper(appKeepers.OVMKeeper)
	appKeepers.MarketKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)

	appKeepers.BetKeeper = betmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[betmoduletypes.StoreKey],
		appKeepers.keys[betmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(betmoduletypes.ModuleName),
		govModAddress,
	)
	appKeepers.BetKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
	appKeepers.BetKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)
	appKeepers.BetKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

	appKeepers.OrderbookKeeper.SetBetKeeper(appKeepers.BetKeeper)
	appKeepers.OrderbookKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
	appKeepers.OrderbookKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

	appKeepers.HouseKeeper = housemodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[housemoduletypes.StoreKey],
		appKeepers.OrderbookKeeper,
		appKeepers.OVMKeeper,
		appKeepers.GetSubspace(housemoduletypes.ModuleName),
		housemodulekeeper.SdkExpectedKeepers{
			AuthzKeeper: appKeepers.AuthzKeeper,
		},
		govModAddress,
	)
	appKeepers.OrderbookKeeper.SetHouseKeeper(appKeepers.HouseKeeper)

	appKeepers.SubaccountKeeper = subaccountmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[subaccountmoduletypes.StoreKey],
		appKeepers.GetSubspace(subaccountmoduletypes.ModuleName),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.OVMKeeper,
		appKeepers.BetKeeper,
		appKeepers.OrderbookKeeper,
		appKeepers.HouseKeeper,
		govModAddress,
	)

	appKeepers.RewardKeeper = rewardmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[rewardmoduletypes.StoreKey],
		appKeepers.keys[rewardmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(rewardmoduletypes.ModuleName),
		appKeepers.BetKeeper,
		appKeepers.OVMKeeper,
		appKeepers.SubaccountKeeper,
		rewardmodulekeeper.SdkExpectedKeepers{
			AuthzKeeper:   appKeepers.AuthzKeeper,
			BankKeeper:    appKeepers.BankKeeper,
			AccountKeeper: appKeepers.AccountKeeper,
		},
		govModAddress,
	)

	// ** Hooks ** \\

	appKeepers.OrderbookKeeper.SetHooks(
		orderbookmoduletypes.NewMultiOrderBookHooks(
			appKeepers.SubaccountKeeper.Hooks(),
		),
	)

	// // SGE modules \\\\

	appKeepers.BetModule = betmodule.NewAppModule(
		appCodec,
		*appKeepers.BetKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.MarketKeeper,
		appKeepers.OrderbookKeeper,
		appKeepers.OVMKeeper,
	)
	appKeepers.MarketModule = marketmodule.NewAppModule(
		appCodec,
		*appKeepers.MarketKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.OVMKeeper,
	)
	appKeepers.HouseModule = housemodule.NewAppModule(
		appCodec,
		appKeepers.HouseKeeper,
	)
	appKeepers.OrderbookModule = orderbookmodule.NewAppModule(
		appCodec,
		*appKeepers.OrderbookKeeper,
	)
	appKeepers.OVMModule = ovmmodule.NewAppModule(
		appCodec,
		*appKeepers.OVMKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)
	appKeepers.RewardModule = rewardmodule.NewAppModule(
		appCodec,
		*appKeepers.RewardKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)

	appKeepers.SubaccountModule = subaccountmodule.NewAppModule(*appKeepers.SubaccountKeeper)

	// wire up x/wasm to IBC
	ibcRouter.AddRoute(wasmtypes.ModuleName, wasm.NewIBCHandler(appKeepers.WasmKeeper, appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCKeeper.ChannelKeeper))

	// Seal the router
	appKeepers.IBCKeeper.SetRouter(ibcRouter)

	// register the proposal types
	govRouter := govtypesv1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypesv1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, sdkparams.NewParamChangeProposalHandler(*appKeepers.ParamsKeeper))

	govConfig := govtypes.DefaultConfig()
	// Set the maximum metadata length for government-related configurations to 10,200, deviating from the default value of 256.
	govConfig.MaxMetadataLen = 10200
	govKeeper := govkeeper.NewKeeper(
		appCodec, runtime.NewKVStoreService(appKeepers.keys[govtypes.StoreKey]),
		appKeepers.AccountKeeper, appKeepers.BankKeeper, appKeepers.StakingKeeper, appKeepers.DistrKeeper, bApp.MsgServiceRouter(),
		govConfig, authtypes.NewModuleAddress(govtypes.ModuleName).String())
	appKeepers.GovKeeper = govKeeper
	appKeepers.GovKeeper.SetLegacyRouter(govRouter)
}

// initParamsKeeper init params keeper and its subspaces.
func (appKeepers *AppKeepers) initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)

	// register the key tables for legacy param subspaces
	keyTable := ibcclienttypes.ParamKeyTable()
	keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})
	paramsKeeper.Subspace(ibchost.ModuleName).WithKeyTable(keyTable)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName).WithKeyTable(ibctransfertypes.ParamKeyTable())
	paramsKeeper.Subspace(icahosttypes.SubModuleName).WithKeyTable(icahosttypes.ParamKeyTable())
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName).WithKeyTable(icacontrollertypes.ParamKeyTable())
	paramsKeeper.Subspace(icqtypes.ModuleName)
	paramsKeeper.Subspace(packetforwardtypes.ModuleName).WithKeyTable(packetforwardtypes.ParamKeyTable())

	paramsKeeper.Subspace(wasmtypes.ModuleName)

	paramsKeeper.Subspace(betmoduletypes.ModuleName)
	paramsKeeper.Subspace(marketmoduletypes.ModuleName)
	paramsKeeper.Subspace(orderbookmoduletypes.ModuleName)
	paramsKeeper.Subspace(ovmmoduletypes.ModuleName)
	paramsKeeper.Subspace(housemoduletypes.ModuleName)
	paramsKeeper.Subspace(rewardmoduletypes.ModuleName)
	paramsKeeper.Subspace(subaccountmoduletypes.ModuleName)

	return paramsKeeper
}

// func NewAppKeeper(
// 	appCodec codec.Codec,
// 	bApp *baseapp.BaseApp,
// 	cdc *codec.LegacyAmino,
// 	maccPerms map[string][]string,
// 	skipUpgradeHeights map[int64]bool,
// 	homePath string,
// 	invCheckPeriod uint,
// 	wasmOpts []wasmkeeper.Option,
// 	appOpts servertypes.AppOptions,
// ) AppKeepers {
// 	appKeepers := AppKeepers{}
// 	// Set keys KVStoreKey, TransientStoreKey, MemoryStoreKey
// 	appKeepers.GenerateKeys()

// 	appKeepers.ParamsKeeper = initParamsKeeper(
// 		appCodec,
// 		cdc,
// 		appKeepers.keys[paramstypes.StoreKey],
// 		appKeepers.tkeys[paramstypes.TStoreKey],
// 	)

// 	govModAddress := authtypes.NewModuleAddress(govtypes.ModuleName).String()

// 	// set the BaseApp's parameter store
// 	appKeepers.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(appCodec, appKeepers.keys[consensusparamtypes.StoreKey], govModAddress)
// 	bApp.SetParamStore(&appKeepers.ConsensusParamsKeeper)

// 	// add capability keeper and ScopeToModule for ibc module
// 	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[capabilitytypes.StoreKey],
// 		appKeepers.memKeys[capabilitytypes.MemStoreKey],
// 	)

// 	// grant capabilities for the ibc and ibc-transfer modules
// 	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
// 	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
// 	appKeepers.ScopedICAControllerKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
// 	appKeepers.ScopedICAHostKeeper = appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
// 	appKeepers.ScopedWasmKeeper = appKeepers.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)

// 	// add keepers
// 	appKeepers.CrisisKeeper = crisiskeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[crisistypes.StoreKey],
// 		invCheckPeriod,
// 		appKeepers.BankKeeper,
// 		authtypes.FeeCollectorName,
// 		govModAddress,
// 	)

// 	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
// 		appCodec,
// 		appKeepers.keys[authtypes.StoreKey],
// 		authtypes.ProtoBaseAccount,
// 		maccPerms,
// 		AccountAddressPrefix,
// 		govModAddress,
// 	)
// 	appKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
// 		appCodec,
// 		appKeepers.keys[banktypes.StoreKey],
// 		appKeepers.AccountKeeper,
// 		BlockedAddresses(maccPerms),
// 		govModAddress,
// 	)
// 	appKeepers.AuthzKeeper = authzkeeper.NewKeeper(
// 		appKeepers.keys[authzkeeper.StoreKey],
// 		appCodec,
// 		bApp.MsgServiceRouter(),
// 		appKeepers.AccountKeeper,
// 	)

// 	groupConfig := group.DefaultConfig()
// 	appKeepers.GroupKeeper = groupkeeper.NewKeeper(
// 		appKeepers.keys[group.StoreKey],
// 		appCodec,
// 		bApp.MsgServiceRouter(),
// 		appKeepers.AccountKeeper, groupConfig,
// 	)

// 	appKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[feegrant.StoreKey],
// 		appKeepers.AccountKeeper,
// 	)

// 	appKeepers.StakingKeeper = stakingkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[stakingtypes.StoreKey],
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		govModAddress,
// 	)

// 	appKeepers.MintKeeper = *mintkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[minttypes.StoreKey],
// 		appKeepers.GetSubspace(minttypes.ModuleName),
// 		appKeepers.AccountKeeper,
// 		mintkeeper.ExpectedKeepers{
// 			StakingKeeper: appKeepers.StakingKeeper,
// 			BankKeeper:    appKeepers.BankKeeper,
// 		},
// 		authtypes.FeeCollectorName,
// 		govModAddress,
// 	)

// 	appKeepers.DistrKeeper = distrkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[distrtypes.StoreKey],
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.StakingKeeper,
// 		authtypes.FeeCollectorName,
// 		govModAddress,
// 	)
// 	appKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
// 		appCodec,
// 		cdc,
// 		appKeepers.keys[slashingtypes.StoreKey],
// 		appKeepers.StakingKeeper,
// 		govModAddress,
// 	)

// 	// register the staking hooks
// 	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
// 	appKeepers.StakingKeeper.SetHooks(
// 		stakingtypes.NewMultiStakingHooks(
// 			appKeepers.DistrKeeper.Hooks(),
// 			appKeepers.SlashingKeeper.Hooks(),
// 		),
// 	)

// 	// UpgradeKeeper must be created before IBCKeeper
// 	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(
// 		skipUpgradeHeights,
// 		appKeepers.keys[upgradetypes.StoreKey],
// 		appCodec,
// 		homePath,
// 		bApp,
// 		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
// 	)

// 	// Create IBC Keeper
// 	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[ibcexported.StoreKey],
// 		appKeepers.GetSubspace(ibcexported.ModuleName),
// 		appKeepers.StakingKeeper,
// 		appKeepers.UpgradeKeeper,
// 		appKeepers.ScopedIBCKeeper,
// 	)

// 	// register the proposal types
// 	govRouter := govv1beta1.NewRouter()
// 	govRouter.
// 		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
// 		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(appKeepers.ParamsKeeper)).
// 		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(appKeepers.UpgradeKeeper)).
// 		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper))

// 	govConfig := govtypes.DefaultConfig()

// 	appKeepers.GovKeeper = govkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[govtypes.StoreKey],
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.StakingKeeper,
// 		bApp.MsgServiceRouter(),
// 		govConfig,
// 		govModAddress,
// 	)
// 	appKeepers.GovKeeper.SetLegacyRouter(govRouter)

// 	// Configure the hooks keeper
// 	hooksKeeper := ibchookskeeper.NewKeeper(
// 		appKeepers.keys[ibchookstypes.StoreKey],
// 	)
// 	appKeepers.IBCHooksKeeper = &hooksKeeper

// 	sgePrefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
// 	wasmHooks := ibc_hooks.NewWasmHooks(appKeepers.IBCHooksKeeper, &appKeepers.WasmKeeper, sgePrefix) // The contract keeper needs to be set later
// 	appKeepers.Ics20WasmHooks = &wasmHooks
// 	appKeepers.HooksICS4Wrapper = ibc_hooks.NewICS4Middleware(
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		appKeepers.Ics20WasmHooks,
// 	)

// 	// IBC Fee Module keeper
// 	appKeepers.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[ibcfeetypes.StoreKey],
// 		appKeepers.IBCKeeper.ChannelKeeper, // more middlewares can be added in future
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		&appKeepers.IBCKeeper.PortKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 	)

// 	appKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[ibctransfertypes.StoreKey],
// 		appKeepers.GetSubspace(ibctransfertypes.ModuleName),
// 		appKeepers.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		&appKeepers.IBCKeeper.PortKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.ScopedTransferKeeper,
// 	)

// 	appKeepers.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[icacontrollertypes.StoreKey],
// 		appKeepers.GetSubspace(icacontrollertypes.SubModuleName),
// 		appKeepers.IBCFeeKeeper,
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		&appKeepers.IBCKeeper.PortKeeper,
// 		appKeepers.ScopedICAControllerKeeper,
// 		bApp.MsgServiceRouter(),
// 	)

// 	appKeepers.ICAHostKeeper = icahostkeeper.NewKeeper(
// 		appCodec, appKeepers.keys[icahosttypes.StoreKey],
// 		appKeepers.GetSubspace(icahosttypes.SubModuleName),
// 		appKeepers.IBCFeeKeeper,
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		&appKeepers.IBCKeeper.PortKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.ScopedICAHostKeeper,
// 		bApp.MsgServiceRouter(),
// 	)

// 	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
// 	appKeepers.EvidenceKeeper = *evidencekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[evidencetypes.StoreKey],
// 		appKeepers.StakingKeeper,
// 		appKeepers.SlashingKeeper,
// 	)

// 	wasmDir := filepath.Join(homePath, "cwvm")
// 	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
// 	if err != nil {
// 		panic("error while reading wasm config: " + err.Error())
// 	}

// 	ibcWasmConfig := ibcwasmtypes.WasmConfig{
// 		DataDir:               filepath.Join(wasmDir, "ibc_08-wasm"),
// 		SupportedCapabilities: "iterator,stargate,abort",
// 		ContractDebugMode:     false,
// 	}

// 	wasmCapabilities := "iterator,staking,stargate,cosmwasm_1_1"

// 	appKeepers.WasmKeeper = wasmkeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[wasmtypes.StoreKey],
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.StakingKeeper,
// 		distrkeeper.NewQuerier(appKeepers.DistrKeeper),
// 		appKeepers.IBCFeeKeeper,
// 		appKeepers.IBCKeeper.ChannelKeeper,
// 		&appKeepers.IBCKeeper.PortKeeper,
// 		appKeepers.ScopedWasmKeeper,
// 		appKeepers.TransferKeeper,
// 		bApp.MsgServiceRouter(),
// 		bApp.GRPCQueryRouter(),
// 		wasmDir,
// 		wasmConfig,
// 		wasmCapabilities,
// 		govModAddress,
// 		wasmOpts...,
// 	)

// 	appKeepers.WasmClientKeeper = ibcwasmkeeper.NewKeeperWithConfig(
// 		appCodec,
// 		appKeepers.keys[ibcwasmtypes.StoreKey],
// 		appKeepers.IBCKeeper.ClientKeeper,
// 		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
// 		ibcWasmConfig,
// 		bApp.GRPCQueryRouter(),
// 	)

// 	// set the contract keeper for the Ics20WasmHooks
// 	appKeepers.ContractKeeper = wasmkeeper.NewDefaultPermissionKeeper(&appKeepers.WasmKeeper)
// 	appKeepers.Ics20WasmHooks.ContractKeeper = &appKeepers.WasmKeeper

// 	// // SGE keepers \\\\

// 	appKeepers.OrderbookKeeper = orderbookmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[orderbookmoduletypes.StoreKey],
// 		appKeepers.GetSubspace(orderbookmoduletypes.ModuleName),
// 		orderbookmodulekeeper.SdkExpectedKeepers{
// 			BankKeeper:     appKeepers.BankKeeper,
// 			AccountKeeper:  appKeepers.AccountKeeper,
// 			FeeGrantKeeper: appKeepers.FeeGrantKeeper,
// 		},
// 		govModAddress,
// 	)

// 	appKeepers.OVMKeeper = ovmmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[ovmmoduletypes.StoreKey],
// 		appKeepers.keys[ovmmoduletypes.MemStoreKey],
// 		appKeepers.GetSubspace(ovmmoduletypes.ModuleName),
// 		govModAddress,
// 	)

// 	appKeepers.MarketKeeper = marketmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[marketmoduletypes.StoreKey],
// 		appKeepers.keys[marketmoduletypes.MemStoreKey],
// 		appKeepers.GetSubspace(marketmoduletypes.ModuleName),
// 		govModAddress,
// 	)
// 	appKeepers.MarketKeeper.SetOVMKeeper(appKeepers.OVMKeeper)
// 	appKeepers.MarketKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)

// 	appKeepers.BetKeeper = betmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[betmoduletypes.StoreKey],
// 		appKeepers.keys[betmoduletypes.MemStoreKey],
// 		appKeepers.GetSubspace(betmoduletypes.ModuleName),
// 		govModAddress,
// 	)
// 	appKeepers.BetKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
// 	appKeepers.BetKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)
// 	appKeepers.BetKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

// 	appKeepers.OrderbookKeeper.SetBetKeeper(appKeepers.BetKeeper)
// 	appKeepers.OrderbookKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
// 	appKeepers.OrderbookKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

// 	appKeepers.HouseKeeper = housemodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[housemoduletypes.StoreKey],
// 		appKeepers.OrderbookKeeper,
// 		appKeepers.OVMKeeper,
// 		appKeepers.GetSubspace(housemoduletypes.ModuleName),
// 		housemodulekeeper.SdkExpectedKeepers{
// 			AuthzKeeper: appKeepers.AuthzKeeper,
// 		},
// 		govModAddress,
// 	)
// 	appKeepers.OrderbookKeeper.SetHouseKeeper(appKeepers.HouseKeeper)

// 	appKeepers.SubaccountKeeper = subaccountmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[subaccountmoduletypes.StoreKey],
// 		appKeepers.GetSubspace(subaccountmoduletypes.ModuleName),
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.OVMKeeper,
// 		appKeepers.BetKeeper,
// 		appKeepers.OrderbookKeeper,
// 		appKeepers.HouseKeeper,
// 		govModAddress,
// 	)

// 	appKeepers.RewardKeeper = rewardmodulekeeper.NewKeeper(
// 		appCodec,
// 		appKeepers.keys[rewardmoduletypes.StoreKey],
// 		appKeepers.keys[rewardmoduletypes.MemStoreKey],
// 		appKeepers.GetSubspace(rewardmoduletypes.ModuleName),
// 		appKeepers.BetKeeper,
// 		appKeepers.OVMKeeper,
// 		appKeepers.SubaccountKeeper,
// 		rewardmodulekeeper.SdkExpectedKeepers{
// 			AuthzKeeper:   appKeepers.AuthzKeeper,
// 			BankKeeper:    appKeepers.BankKeeper,
// 			AccountKeeper: appKeepers.AccountKeeper,
// 		},
// 		govModAddress,
// 	)

// 	// ** Hooks ** \\

// 	appKeepers.OrderbookKeeper.SetHooks(
// 		orderbookmoduletypes.NewMultiOrderBookHooks(
// 			appKeepers.SubaccountKeeper.Hooks(),
// 		),
// 	)

// 	// // SGE modules \\\\

// 	appKeepers.BetModule = betmodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.BetKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.MarketKeeper,
// 		appKeepers.OrderbookKeeper,
// 		appKeepers.OVMKeeper,
// 	)
// 	appKeepers.MarketModule = marketmodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.MarketKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 		appKeepers.OVMKeeper,
// 	)
// 	appKeepers.HouseModule = housemodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.HouseKeeper,
// 	)
// 	appKeepers.OrderbookModule = orderbookmodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.OrderbookKeeper,
// 	)
// 	appKeepers.OVMModule = ovmmodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.OVMKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 	)
// 	appKeepers.RewardModule = rewardmodule.NewAppModule(
// 		appCodec,
// 		*appKeepers.RewardKeeper,
// 		appKeepers.AccountKeeper,
// 		appKeepers.BankKeeper,
// 	)

// 	appKeepers.SubaccountModule = subaccountmodule.NewAppModule(*appKeepers.SubaccountKeeper)

// 	//// IBC modules \\\\
// 	appKeepers.IBCModule = ibc.NewAppModule(appKeepers.IBCKeeper)
// 	appKeepers.IBCFeeModule = ibcfee.NewAppModule(appKeepers.IBCFeeKeeper)
// 	appKeepers.TransferModule = transfer.NewAppModule(appKeepers.TransferKeeper)
// 	appKeepers.ICAModule = ica.NewAppModule(&appKeepers.ICAControllerKeeper, &appKeepers.ICAHostKeeper)

// 	// IBC stacks \\\
// 	var transferStack ibcporttypes.IBCModule
// 	transferStack = transfer.NewIBCModule(appKeepers.TransferKeeper)
// 	transferStack = ibcfee.NewIBCMiddleware(transferStack, appKeepers.IBCFeeKeeper)

// 	var icaControllerStack ibcporttypes.IBCModule
// 	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, appKeepers.ICAControllerKeeper)
// 	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, appKeepers.IBCFeeKeeper)

// 	var icaHostStack ibcporttypes.IBCModule
// 	icaHostStack = icahost.NewIBCModule(appKeepers.ICAHostKeeper)
// 	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, appKeepers.IBCFeeKeeper)

// 	// Create fee enabled wasm ibc Stack
// 	var wasmStack ibcporttypes.IBCModule
// 	wasmStack = wasm.NewIBCHandler(appKeepers.WasmKeeper, appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCFeeKeeper)
// 	wasmStack = ibcfee.NewIBCMiddleware(wasmStack, appKeepers.IBCFeeKeeper)

// 	// Create static IBC router, add transfer route, then set and seal it
// 	ibcRouter := ibcporttypes.NewRouter()
// 	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerStack)
// 	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostStack)
// 	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
// 	ibcRouter.AddRoute(wasmtypes.ModuleName, wasmStack)

// 	appKeepers.IBCKeeper.SetRouter(ibcRouter)

// 	/****  Module Options ****/
// 	return appKeepers
// }

// GetSubspace returns a param subspace for a given module name.
func (appKeepers *AppKeepers) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := appKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(betmoduletypes.ModuleName)
	paramsKeeper.Subspace(marketmoduletypes.ModuleName)
	paramsKeeper.Subspace(orderbookmoduletypes.ModuleName)
	paramsKeeper.Subspace(ovmmoduletypes.ModuleName)
	paramsKeeper.Subspace(housemoduletypes.ModuleName)
	paramsKeeper.Subspace(rewardmoduletypes.ModuleName)
	paramsKeeper.Subspace(subaccountmoduletypes.ModuleName)
	paramsKeeper.Subspace(wasmtypes.ModuleName)

	return paramsKeeper
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms(maccPerms map[string][]string) map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}
