package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/sge-network/sge/x/reward/types"
)

func CmdCreateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-campaign [uid] [ticket]",
		Short: "Create a new campaign",
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

			msg := types.NewMsgCreateCampaign(
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

func CmdUpdateCampaign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-campaign [uid] [ticket]",
		Short: "Update a campaign",
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

			msg := types.NewMsgUpdateCampaign(
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