package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/auction/types"
)

func CmdCreateBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bid [recType] [auctionId] [bidNo] [itemId] [buyerId] [bidPrice] [bidTime]",
		Short: "Creates a new bid",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsAuctionId := string(args[1])
			argsBidNo := string(args[2])
			argsItemId := string(args[3])
			argsBuyerId := string(args[4])
			argsBidPrice := string(args[5])
			argsBidTime := string(args[6])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBid(clientCtx.GetFromAddress().String(), string(argsRecType), string(argsAuctionId), string(argsBidNo), string(argsItemId), string(argsBuyerId), string(argsBidPrice), string(argsBidTime))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bid [id] [recType] [auctionId] [bidNo] [itemId] [buyerId] [bidPrice] [bidTime]",
		Short: "Update a bid",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsAuctionId := string(args[2])
			argsBidNo := string(args[3])
			argsItemId := string(args[4])
			argsBuyerId := string(args[5])
			argsBidPrice := string(args[6])
			argsBidTime := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBid(clientCtx.GetFromAddress().String(), id, string(argsRecType), string(argsAuctionId), string(argsBidNo), string(argsItemId), string(argsBuyerId), string(argsBidPrice), string(argsBidTime))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-bid [id]",
		Short: "Delete a bid by id",
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

			msg := types.NewMsgDeleteBid(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
