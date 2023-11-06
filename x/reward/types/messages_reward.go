package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrtypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sge-network/sge/utils"
)

const (
	TypeMsgGrantReward = "apply_reward"
)

var _ sdk.Msg = &MsgGrantReward{}

func NewMsgGrantReward(
	creator string,
	campaignUID string,
	ticket string,
) *MsgGrantReward {
	return &MsgGrantReward{
		Creator:     creator,
		CampaignUid: campaignUID,
		Ticket:      ticket,
	}
}

func (msg *MsgGrantReward) Route() string {
	return RouterKey
}

func (msg *MsgGrantReward) Type() string {
	return TypeMsgGrantReward
}

func (msg *MsgGrantReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGrantReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGrantReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !utils.IsValidUID(msg.CampaignUid) {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidRequest, "invalid campaign uid")
	}

	if msg.Ticket == "" {
		return sdkerrors.Wrapf(sdkerrtypes.ErrInvalidRequest, "invalid ticket")
	}

	return nil
}

// EmitEvent emits the event for the message success.
func (msg *MsgGrantReward) EmitEvent(ctx *sdk.Context, campaignUID string, recevier Receiver) {
	emitter := utils.NewEventEmitter(ctx, attributeValueCategory)
	emitter.AddMsg(TypeMsgGrantReward, msg.Creator,
		sdk.NewAttribute(attributeKeyCampaignUID, campaignUID),
		sdk.NewAttribute(attributeKeyDistributions, recevier.String()),
	)
	emitter.Emit()
}
