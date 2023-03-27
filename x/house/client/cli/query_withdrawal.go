package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/sge-network/sge/x/house/types"
	"github.com/spf13/cobra"
)

// GetCmdQueryWithdrawalsByAcc implements the command to query all the withdrawals made from one account address.
func GetCmdQueryWithdrawalsByAcc() *cobra.Command {
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	cmd := &cobra.Command{
		Use:   "withdrawals-by-acc [acc]",
		Short: "Query all withdrawals made by one acc",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query withdrawals for an individual acc on all houses.

Example:
$ %s query house withdrawals-by-acc %s1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p
`,
				version.AppName, bech32PrefixAccAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			depAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryWithdrawalsByAccRequest{
				Address:    depAddr.String(),
				Pagination: pageReq,
			}

			res, err := queryClient.WithdrawalsByAcc(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "withdrawals")

	return cmd
}
