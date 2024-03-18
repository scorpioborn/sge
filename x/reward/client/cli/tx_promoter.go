package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/sge-network/sge/x/reward/types"
)

func CmdSetPromoterConf() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-promoter-conf [uid] [ticket]",
		Short: "Set promoter config",
		Long:  "Set promoter config of a promoter by uid and the ticket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			argUID := args[0]

			// Get value arguments
			argTicket := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPromoterConfig(
				clientCtx.GetFromAddress().String(),
				argUID,
				argTicket,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
