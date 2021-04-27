package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/trans/types"
)

func CmdCreateTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-transaction [recType] [auctionId] [itemId] [transType] [userId] [transDate] [hammerTime] [hammerPrice] [details] [status]",
		Short: "Creates a new transaction",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsAuctionId := string(args[1])
			argsItemId := string(args[2])
			argsTransType := string(args[3])
			argsUserId := string(args[4])
			argsTransDate := string(args[5])
			argsHammerTime := string(args[6])
			argsHammerPrice := string(args[7])
			argsDetails := string(args[8])
			argsStatus := string(args[9])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTransaction(clientCtx.GetFromAddress().String(), string(argsRecType), string(argsAuctionId), string(argsItemId), string(argsTransType), string(argsUserId), string(argsTransDate), string(argsHammerTime), string(argsHammerPrice), string(argsDetails), string(argsStatus))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-transaction [id] [recType] [auctionId] [itemId] [transType] [userId] [transDate] [hammerTime] [hammerPrice] [details] [status]",
		Short: "Update a transaction",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsAuctionId := string(args[2])
			argsItemId := string(args[3])
			argsTransType := string(args[4])
			argsUserId := string(args[5])
			argsTransDate := string(args[6])
			argsHammerTime := string(args[7])
			argsHammerPrice := string(args[8])
			argsDetails := string(args[9])
			argsStatus := string(args[10])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTransaction(clientCtx.GetFromAddress().String(), id, string(argsRecType), string(argsAuctionId), string(argsItemId), string(argsTransType), string(argsUserId), string(argsTransDate), string(argsHammerTime), string(argsHammerPrice), string(argsDetails), string(argsStatus))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-transaction [id] [recType] [auctionId] [itemId] [transType] [userId] [transDate] [hammerTime] [hammerPrice] [details] [status]",
		Short: "Delete a transaction by id",
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

			msg := types.NewMsgDeleteTransaction(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
