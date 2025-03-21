package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/market interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgAdd{}, "market/Add")
	legacy.RegisterAminoMsg(cdc, &MsgResolve{}, "market/Resolve")
	legacy.RegisterAminoMsg(cdc, &MsgUpdate{}, "market/Update")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "market/MsgUpdateParams")
}

// RegisterInterfaces registers the module interface types
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil))
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAdd{},
		&MsgResolve{},
		&MsgUpdate{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// Amino is the legacy Amino codec
var Amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(Amino)
	// cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
