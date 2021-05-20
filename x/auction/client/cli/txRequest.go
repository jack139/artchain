package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/auction/types"
)

func CmdCreateRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-request [recType] [itemId] [auctionHouseId] [SellerId] [requestDate] [reservePrice] [status] [openDate] [closeDate]",
		Short: "Creates a new request",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsItemId := string(args[1])
			argsAuctionHouseId := string(args[2])
			argsSellerId := string(args[3])
			argsRequestDate := string(args[4])
			argsReservePrice := string(args[5])
			argsStatus := string(args[6])
			argsOpenDate := string(args[7])
			argsCloseDate := string(args[8])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRequest(clientCtx.GetFromAddress().String(), string(argsRecType), 
				string(argsItemId), string(argsAuctionHouseId), string(argsSellerId), string(argsRequestDate), 
				string(argsReservePrice), string(argsStatus), string(argsOpenDate), string(argsCloseDate), "")
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-request [id] [recType] [itemId] [auctionHouseId] [SellerId] [requestDate] [reservePrice] [status] [openDate] [closeDate]",
		Short: "Update a request",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsItemId := string(args[2])
			argsAuctionHouseId := string(args[3])
			argsSellerId := string(args[4])
			argsRequestDate := string(args[5])
			argsReservePrice := string(args[6])
			argsStatus := string(args[7])
			argsOpenDate := string(args[8])
			argsCloseDate := string(args[9])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRequest(clientCtx.GetFromAddress().String(), id, string(argsRecType), 
				string(argsItemId), string(argsAuctionHouseId), string(argsSellerId), string(argsRequestDate), 
				string(argsReservePrice), string(argsStatus), string(argsOpenDate), string(argsCloseDate), "")
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-request [id]",
		Short: "Delete a request by id",
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

			msg := types.NewMsgDeleteRequest(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
