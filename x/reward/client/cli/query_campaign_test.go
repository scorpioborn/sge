package cli_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tmcli "github.com/tendermint/tendermint/libs/cli"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/sge-network/sge/testutil/network"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/testutil/sample"
	"github.com/sge-network/sge/x/reward/client/cli"
	"github.com/sge-network/sge/x/reward/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithCampaignObjects(t *testing.T, n int) (*network.Network, []types.Campaign) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		campaign := types.Campaign{
			UID:           uuid.NewString(),
			Creator:       sample.AccAddress(),
			FunderAddress: sample.AccAddress(),
			StartTS:       uint64(time.Now().Unix()),
			EndTS:         uint64(time.Now().Add(5 * time.Minute).Unix()),
			RewardType:    types.RewardType_REWARD_TYPE_AFFILIATION,
			RewardDefs: []types.Definition{{
				RecType:    types.ReceiverType_RECEIVER_TYPE_REFEREE,
				Amount:     sdkmath.NewInt(100),
				RecAccType: types.ReceiverAccType_RECEIVER_ACC_TYPE_MAIN,
			}},
			Pool: types.Pool{Spent: sdkmath.NewInt(100), Total: sdkmath.NewInt(1000)},
		}

		nullify.Fill(&campaign)
		state.CampaignList = append(state.CampaignList, campaign)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.CampaignList
}

func TestShowCampaign(t *testing.T) {
	net, objs := networkWithCampaignObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.Campaign
	}{
		{
			desc:    "found",
			idIndex: objs[0].UID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCampaign(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryCampaignResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Campaign)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Campaign),
				)
			}
		})
	}
}

func TestListCampaign(t *testing.T) {
	net, objs := networkWithCampaignObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
			require.NoError(t, err)
			var resp types.QueryCampaignAllResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Campaign),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
			require.NoError(t, err)
			var resp types.QueryCampaignAllResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Campaign),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaign(), args)
		require.NoError(t, err)
		var resp types.QueryCampaignAllResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Campaign),
		)
	})
}