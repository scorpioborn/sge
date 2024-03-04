package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrtypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sge-network/sge/utils"
)

const (
	TypeMsgSetPromoterConf = "set_promoter_conf"
)

var _ sdk.Msg = &MsgSetPromoterConf{}

func NewMsgSetPromoterConfig(
	creator string,
	uid string,
	ticket string,
) *MsgSetPromoterConf {
	return &MsgSetPromoterConf{
		Creator: creator,
		Uid:     uid,
		Ticket:  ticket,
	}
}

func (msg *MsgSetPromoterConf) Route() string {
	return RouterKey
}

func (msg *MsgSetPromoterConf) Type() string {
	return TypeMsgSetPromoterConf
}

func (msg *MsgSetPromoterConf) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPromoterConf) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPromoterConf) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !utils.IsValidUID(msg.Uid) {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidRequest, "invalid uid")
	}

	if msg.Ticket == "" {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidRequest, "invalid ticket")
	}

	return nil
}

// EmitEvent emits the event for the message success.
func (msg *MsgSetPromoterConf) EmitEvent(ctx *sdk.Context, conf PromoterConf) {
	emitter := utils.NewEventEmitter(ctx, attributeValueCategory)
	emitter.AddMsg(TypeMsgSetPromoterConf, msg.Creator,
		sdk.NewAttribute(attributeKeyUID, msg.Uid),
		sdk.NewAttribute(attributeKeyPromoterConf, conf.String()),
	)
	emitter.Emit()
}
