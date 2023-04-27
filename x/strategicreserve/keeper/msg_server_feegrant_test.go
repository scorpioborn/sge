package keeper_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sge-network/sge/testutil/sample"
	simappUtil "github.com/sge-network/sge/testutil/simapp"
	"github.com/sge-network/sge/x/strategicreserve/types"
	"github.com/stretchr/testify/require"
)

func TestMsgServerInvokeFeeGrant(t *testing.T) {
	tApp, k, msgk, ctx, wctx := setupMsgServerAndApp(t)

	granter := k.GetModuleAddress(types.SRPoolName)
	grantee := simappUtil.TestParamUsers["user1"]

	t.Run("valid feegrant invoke", func(t *testing.T) {
		validEmptyTicketClaims := jwt.MapClaims{
			"grantee": grantee.Address.String(),
			"exp":     9999999999,
			"iat":     1111111111,
		}
		validEmptyTicket, err := createJwtTicket(validEmptyTicketClaims)
		require.NoError(t, err)

		response, err := msgk.InvokeFeeGrant(wctx, types.NewMsgInvokeFeeGrant(sample.AccAddress(), validEmptyTicket))
		require.NoError(t, err)
		require.NotNil(t, response)

		feeGrant, found := k.GetFeeGrant(ctx, grantee.Address.String())
		require.True(t, found)
		require.Equal(t, grantee.Address.String(), feeGrant.Grantee)

		allowance, err := tApp.FeeGrantKeeper.GetAllowance(ctx, granter, grantee.Address)
		require.NoError(t, err)
		require.NotNil(t, allowance)
	})
}

func TestMsgServerRevokeFeeGrant(t *testing.T) {
	k, msgk, ctx, wctx := setupMsgServerAndKeeper(t)

	ctx = ctx.WithBlockTime(time.Now())

	grantee := simappUtil.TestParamUsers["user1"]

	k.SetFeeGrant(ctx, types.NewFeeGrant(sample.AccAddress(), grantee.Address.String(), ctx.BlockTime().Unix()))

	t.Run("valid feegrant revoke", func(t *testing.T) {
		validEmptyTicketClaims := jwt.MapClaims{
			"grantee": grantee.Address.String(),
			"exp":     9999999999,
			"iat":     1111111111,
		}
		validEmptyTicket, err := createJwtTicket(validEmptyTicketClaims)
		require.NoError(t, err)

		response, err := msgk.RevokeFeeGrant(wctx, types.NewMsgRevokeFeeGrant(sample.AccAddress(), validEmptyTicket))
		require.NoError(t, err)
		require.NotNil(t, response)

		feeGrant, found := k.GetFeeGrant(ctx, grantee.Address.String())
		require.False(t, found)
		require.Equal(t, feeGrant.Creator, "")
	})
}

func createJwtTicket(claim jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claim)
	return token.SignedString(simappUtil.TestOVMPrivateKeys[0])
}
