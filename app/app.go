package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/posthandler"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibcwasmkeeper "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/keeper"
	ibcwasmtypes "github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	"github.com/sge-network/sge/app/keepers"
	sgeappparams "github.com/sge-network/sge/app/params"
	"github.com/sge-network/sge/app/upgrades"
	v10 "github.com/sge-network/sge/app/upgrades/v10"
	v9 "github.com/sge-network/sge/app/upgrades/v9"
	// unnamed import of statik for swagger UI support
	// TODO: Check if needed
	// _ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = moduleAccountPermissions

	Upgrades = []upgrades.Upgrade{
		v9.Upgrade,
		v10.Upgrade,
	}
	// default wasmd binary maximum allowed size
	defaultMaxWasmSize = 1500 * 1024 // 1.5 MB
)

var (
	_ runtime.AppI            = (*SgeApp)(nil)
	_ servertypes.Application = (*SgeApp)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+appName)
}

// SgeApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SgeApp struct {
	*baseapp.BaseApp
	keepers.AppKeepers

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// the module manager
	mm *module.Manager

	// basics manager
	ModuleBasics module.BasicManager

	// simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
	homePath     string
}

// NewSgeApp returns a reference to an initialized Sge.
func NewSgeApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig sgeappparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	wasmOpts []wasmkeeper.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *SgeApp {
	wasmtypes.MaxWasmSize = defaultMaxWasmSize
	wasmtypes.MaxProposalWasmSize = wasmtypes.MaxWasmSize

	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	app := &SgeApp{
		AppKeepers:        keepers.AppKeepers{},
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
	}

	app.homePath = homePath
	dataDir := filepath.Join(homePath, "data")
	wasmDir := filepath.Join(homePath, "wasm")
	ibcWasmConfig := ibcwasmtypes.WasmConfig{
		DataDir:               filepath.Join(homePath, "ibc_08-wasm"),
		SupportedCapabilities: []string{"iterator", "stargate", "abort"},
		ContractDebugMode:     false,
	}
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	// Uncomment this for debugging contracts. In the future this could be made into a param passed by the tests
	// wasmConfig.ContractDebugMode = true
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// Setup keepers
	app.InitSpecialKeepers(
		appCodec,
		bApp,
		wasmDir,
		cdc,
		invCheckPeriod,
		skipUpgradeHeights,
		homePath,
	)
	app.setupUpgradeStoreLoaders()
	app.InitNormalKeepers(
		appCodec,
		encodingConfig,
		bApp,
		maccPerms,
		dataDir,
		wasmDir,
		wasmConfig,
		wasmOpts,
		app.BlockedAddresses(),
		ibcWasmConfig,
	)

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(appModules(app, encodingConfig, skipGenesisInvariants)...)

	// Upgrades from v0.50.x onwards happen in pre block
	app.mm.SetOrderPreBlockers(upgradetypes.ModuleName)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(orderBeginBlockers()...)

	app.mm.SetOrderEndBlockers(orderEndBlockers()...)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(orderInitBlockers()...)

	app.mm.RegisterInvariants(app.CrisisKeeper)

	app.configurator = module.NewConfigurator(
		app.appCodec,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
	)
	err = app.mm.RegisterServices(app.configurator)
	if err != nil {
		panic(err)
	}

	// // v47 - no dependecy injection, so register new gRPC services.
	// //#nosec
	// autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))
	// reflectionSvc, err := runtimeservices.NewReflectionService()
	// if err != nil {
	// 	panic(err)
	// }
	// reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	// if err != nil {
	// 	panic("error while reading wasm config: " + err.Error())
	// }

	// Override the gov ModuleBasic with all the custom proposal handers, otherwise we lose them in the CLI.
	app.ModuleBasics = module.NewBasicManagerFromManager(
		app.mm,
		map[string]module.AppModuleBasic{
			"gov": gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
		},
	)

	app.setupUpgradeHandlers()

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		simulationModules(app, encodingConfig, skipGenesisInvariants)...)

	app.sm.RegisterStoreDecoders()

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.mm.Modules))

	reflectionSvc := getReflectionService()
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// initialize stores
	app.MountKVStores(app.GetKVStoreKey())
	app.MountTransientStores(app.GetTransientStoreKey())
	app.MountMemoryStores(app.GetMemoryStoreKey())

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetPreBlocker(app.PreBlocker)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				FeegrantKeeper:  app.FeeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},

			IBCKeeper:             app.IBCKeeper,
			WasmConfig:            &wasmConfig,
			WasmKeeper:            app.WasmKeeper,
			TXCounterStoreService: runtime.NewKVStoreService(app.GetKey(wasmtypes.StoreKey)),
			CircuitKeeper:         app.CircuitBreakerKeeper,
		},
	)
	if err != nil {
		panic(fmt.Errorf(ErrTextFailedToCreateAnteHandler, err))
	}
	app.SetAnteHandler(anteHandler)
	if err := app.setPostHandler(); err != nil {
		panic(fmt.Errorf("%s", err))
	}
	app.SetEndBlocker(app.EndBlocker)
	app.SetPrecommiter(app.Precommitter)
	app.SetPrepareCheckStater(app.PrepareCheckStater)

	if manager := app.SnapshotManager(); manager != nil {
		err = manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), app.AppKeepers.WasmKeeper),
		)
		if err != nil {
			panic("failed to register snapshot extension: " + err.Error())
		}

		err = manager.RegisterExtensions(
			ibcwasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), app.IBCWasmClientKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})

		// https://github.com/cosmos/ibc-go/pull/5439
		if err := ibcwasmkeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("wasmlckeeper failed initialize pinned codes %s", err))
		}

		if err := app.AppKeepers.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("app.AppKeepers.WasmKeeper failed initialize pinned codes %s", err))
		}

		app.AppKeepers.CapabilityKeeper.Seal()
	}

	return app
}

func (app *SgeApp) setPostHandler() error {
	postHandler, err := posthandler.NewPostHandler(
		posthandler.HandlerOptions{},
	)
	if err != nil {
		return err
	}
	app.SetPostHandler(postHandler)
	return nil
}

// Precommitter application updates before the commital of a block after all transactions have been delivered.
func (app *SgeApp) Precommitter(ctx sdk.Context) {
	mm := app.ModuleManager()
	if err := mm.Precommit(ctx); err != nil {
		panic(err)
	}
}

// PreBlocker application updates before each begin block.
func (app *SgeApp) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	// Set gas meter to the free gas meter.
	// This is because there is currently non-deterministic gas usage in the
	// pre-blocker, e.g. due to hydration of in-memory data structures.
	//
	// Note that we don't need to reset the gas meter after the pre-blocker
	// because Go is pass by value.
	ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
	mm := app.ModuleManager()
	return mm.PreBlock(ctx)
}

func (app *SgeApp) PrepareCheckStater(ctx sdk.Context) {
	mm := app.ModuleManager()
	if err := mm.PrepareCheckState(ctx); err != nil {
		panic(err)
	}
}

func (app *SgeApp) ModuleManager() module.Manager {
	return *app.mm
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// Sge. It is useful for tests and clients who do not want to construct the
// full sge application
func MakeCodecs() (codec.Codec,
	*codec.LegacyAmino,
) {
	config := MakeEncodingConfig()
	return config.Marshaler, config.Amino
}

// Name returns the name of the App
func (app *SgeApp) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app *SgeApp) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// BeginBlocker application updates every begin block.
func (app *SgeApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	return app.mm.BeginBlock(ctx)
}

// EndBlocker application updates every end block.
func (app *SgeApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.mm.EndBlock(ctx)
}

// InitChainer application update at chain initialization.
func (app *SgeApp) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	err := app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SgeApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SgeApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)

	//#nosec
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SgeApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SgeApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Sge's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SgeApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Sge's InterfaceRegistry
func (app *SgeApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SgeApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SgeApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SgeApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(),
		clientCtx,
		app.BaseApp.Simulate,
		app.interfaceRegistry,
	)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *SgeApp) RegisterTendermintService(clientCtx client.Context) {
	cmtservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

// RegisterNodeService registers the node gRPC Query service.
func (app *SgeApp) RegisterNodeService(clientCtx client.Context, cfg config.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *SgeApp) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			storeUpgrade := upgrade.StoreUpgrades
			app.SetStoreLoader(
				upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrade),
			)
		}
	}
}

func (app *SgeApp) setupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.mm,
				app.configurator,
				&app.AppKeepers,
			),
		)
	}
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// SimulationManager implements the SimulationApp interface
func (app *SgeApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// BlockedAddresses returns all the app's blocked account addresses.
func (app *SgeApp) BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// we cache the reflectionService to save us time within tests.
var cachedReflectionService *runtimeservices.ReflectionService = nil

func getReflectionService() *runtimeservices.ReflectionService {
	if cachedReflectionService != nil {
		return cachedReflectionService
	}
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	cachedReflectionService = reflectionSvc
	return reflectionSvc
}
