package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sge-network/sge/consts"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/x/orderbook/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestOrderBookQuerySingle(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrderBook(tApp, k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryOrderBookRequest
		response *types.QueryOrderBookResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryOrderBookRequest{
				OrderBookUid: msgs[0].UID,
			},
			response: &types.QueryOrderBookResponse{OrderBook: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryOrderBookRequest{
				OrderBookUid: msgs[1].UID,
			},
			response: &types.QueryOrderBookResponse{OrderBook: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryOrderBookRequest{
				OrderBookUid: cast.ToString(100000),
			},
			err: status.Error(codes.NotFound, "order book 100000 not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.OrderBook(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOrderBooksQueryPaginated(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrderBook(tApp, k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryOrderBooksRequest {
		return &types.QueryOrderBooksRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBooks(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orderbooks), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Orderbooks),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBooks(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Orderbooks), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Orderbooks),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.OrderBooks(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Orderbooks),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.OrderBooks(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest))
	})
}

func TestParticipationQuerySingle(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNParticipation(tApp, k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryOrderBookParticipationRequest
		response *types.QueryOrderBookParticipationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryOrderBookParticipationRequest{
				OrderBookUid:       msgs[0].OrderBookUID,
				ParticipationIndex: msgs[0].Index,
			},
			response: &types.QueryOrderBookParticipationResponse{OrderBookParticipation: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryOrderBookParticipationRequest{
				OrderBookUid:       msgs[1].OrderBookUID,
				ParticipationIndex: msgs[1].Index,
			},
			response: &types.QueryOrderBookParticipationResponse{OrderBookParticipation: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryOrderBookParticipationRequest{
				OrderBookUid:       cast.ToString(100000),
				ParticipationIndex: 100,
			},
			err: status.Error(codes.NotFound, "order book participation 100000, 100 not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.OrderBookParticipation(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOrderBookParticipationsQueryPaginated(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNParticipation(tApp, k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryOrderBookParticipationsRequest {
		return &types.QueryOrderBookParticipationsRequest{
			OrderBookUid: testOrderBookUID,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBookParticipations(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrderBookParticipations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrderBookParticipations),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBookParticipations(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrderBookParticipations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrderBookParticipations),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.OrderBookParticipations(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.OrderBookParticipations),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.OrderBookParticipations(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest))
	})
}

func TestOrderBookEcposureQuerySingle(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrderBookOddsExposure(tApp, k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryOrderBookExposureRequest
		response *types.QueryOrderBookExposureResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryOrderBookExposureRequest{
				OrderBookUid: msgs[0].OrderBookUID,
				OddsUid:      msgs[0].OddsUID,
			},
			response: &types.QueryOrderBookExposureResponse{OrderBookExposure: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryOrderBookExposureRequest{
				OrderBookUid: msgs[1].OrderBookUID,
				OddsUid:      msgs[1].OddsUID,
			},
			response: &types.QueryOrderBookExposureResponse{OrderBookExposure: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryOrderBookExposureRequest{
				OrderBookUid: cast.ToString(100000),
				OddsUid:      "9ca3b818-cfc3-43cd-987b-4cffdf01cc5e",
			},
			err: status.Error(codes.NotFound, "order book exposure 100000, 9ca3b818-cfc3-43cd-987b-4cffdf01cc5e not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.OrderBookExposure(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOrderBookExposuresQueryPaginated(t *testing.T) {
	tApp, k, ctx := setupKeeperAndApp(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOrderBookOddsExposure(tApp, k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryOrderBookExposuresRequest {
		return &types.QueryOrderBookExposuresRequest{
			OrderBookUid: testOrderBookUID,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBookExposures(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrderBookExposures), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrderBookExposures),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.OrderBookExposures(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OrderBookExposures), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OrderBookExposures),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.OrderBookExposures(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.OrderBookExposures),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.OrderBookParticipations(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, consts.ErrTextInvalidRequest))
	})
}
