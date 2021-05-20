package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/inventory/types"
)

func CmdCreateReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-review [recType] [itemId] [reviewerId] [reviewDetail] [reviewDate] [upCount] [downCount]",
		Short: "Creates a new review",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsItemId := string(args[1])
			argsReviewerId := string(args[2])
			argsReviewDetail := string(args[3])
			argsReviewDate := string(args[4])
			argsUpCount := string(args[5])
			argsDownCount := string(args[6])
			argsStatus := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateReview(clientCtx.GetFromAddress().String(), string(argsRecType), 
				string(argsItemId), string(argsReviewerId), string(argsReviewDetail), string(argsReviewDate), 
				string(argsUpCount), string(argsDownCount), string(argsStatus), "")
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-review [id] [recType] [itemId] [reviewerId] [reviewDetail] [reviewDate] [upCount] [downCount]",
		Short: "Update a review",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsItemId := string(args[2])
			argsReviewerId := string(args[3])
			argsReviewDetail := string(args[4])
			argsReviewDate := string(args[5])
			argsUpCount := string(args[6])
			argsDownCount := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateReview(clientCtx.GetFromAddress().String(), id, string(argsRecType), 
				string(argsItemId), string(argsReviewerId), string(argsReviewDetail), string(argsReviewDate), 
				string(argsUpCount), string(argsDownCount), "WAIT","")
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteReview() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-review [id]",
		Short: "Delete a review by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteReview(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
