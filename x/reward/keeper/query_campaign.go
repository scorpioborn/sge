package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sge-network/sge/x/reward/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CampaignAll(goCtx context.Context, req *types.QueryCampaignAllRequest) (*types.QueryCampaignAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var campaigns []types.Campaign
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	campaignStore := prefix.NewStore(store, types.CampaignKeyPrefix)

	pageRes, err := query.Paginate(campaignStore, req.Pagination, func(key []byte, value []byte) error {
		var campaign types.Campaign
		if err := k.cdc.Unmarshal(value, &campaign); err != nil {
			return err
		}

		campaigns = append(campaigns, campaign)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCampaignAllResponse{Campaign: campaigns, Pagination: pageRes}, nil
}

func (k Keeper) Campaign(goCtx context.Context, req *types.QueryCampaignRequest) (*types.QueryCampaignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCampaign(
		ctx,
		req.Uid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryCampaignResponse{Campaign: val}, nil
}
