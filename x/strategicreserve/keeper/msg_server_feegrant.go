package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sge-network/sge/x/strategicreserve/types"
)

// InvokeFeeGrant accepts ticket containing creation feegrant and return response after processing
func (k msgServer) InvokeFeeGrant(goCtx context.Context, msg *types.MsgInvokeFeeGrant) (*types.MsgInvokeFeeGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var payload types.InvokeFeeGrantPayload
	if err := k.ovmKeeper.VerifyTicketUnmarshal(goCtx, msg.Ticket, &payload); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketVerification, "%s", err)
	}

	if err := payload.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketPayloadValidation, "%s", err)
	}

	feeGrant, err := k.Keeper.InvokeFeeGrant(ctx, msg.Creator, payload.Grantee)
	if err != nil {
		return nil, err
	}

	// emitFeeGrantEvent(ctx, types.TypeMsgCreateMarkets, feeGrant.UID, feeGrant.BookUID, msg.Creator)

	return &types.MsgInvokeFeeGrantResponse{
		FeeGrant: &feeGrant,
	}, nil
}

// RevokeFeeGrant accepts ticket containing revoke feegrant and return response after processing.
func (k msgServer) RevokeFeeGrant(goCtx context.Context, msg *types.MsgRevokeFeeGrant) (*types.MsgRevokeFeeGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var payload types.RevokeFeeGrantPayload
	if err := k.ovmKeeper.VerifyTicketUnmarshal(goCtx, msg.Ticket, &payload); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketVerification, "%s", err)
	}

	feeGrant, found := k.Keeper.GetFeeGrant(ctx, payload.Grantee)
	if !found {
		return nil, types.ErrMarketNotFound
	}

	// update market is not valid, return error
	if err := payload.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInTicketPayloadValidation, "%s", err)
	}

	// update market is successful, update the module state
	// the sdk allowance has an expire time, so that will be removed by feegrant module.
	k.Keeper.RemoveFeeGrant(ctx, feeGrant)

	// // emitFeeGrantEvent(ctx, types.TypeMsgCreateMarkets, feeGrant.UID, feeGrant.BookUID, msg.Creator)

	return &types.MsgRevokeFeeGrantResponse{FeeGrant: nil}, nil
}
