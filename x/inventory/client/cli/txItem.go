package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/inventory/types"
)

func CmdCreateItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-item [recType] [itemDesc] [itemDetail] [itemDate] [itemType] [itemSubject] [itemMedia] [itemSize] [itemImage] [AESKey] [itemBasePrice] [currentOwnerId]",
		Short: "Creates a new item",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsItemDesc := string(args[1])
			argsItemDetail := string(args[2])
			argsItemDate := string(args[3])
			argsItemType := string(args[4])
			argsItemSubject := string(args[5])
			argsItemMedia := string(args[6])
			argsItemSize := string(args[7])
			argsItemImage := string(args[8])
			argsAESKey := string(args[9])
			argsItemBasePrice := string(args[10])
			argsCurrentOwnerId := string(args[11])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateItem(clientCtx.GetFromAddress().String(), string(argsRecType), string(argsItemDesc), string(argsItemDetail), string(argsItemDate), string(argsItemType), string(argsItemSubject), string(argsItemMedia), string(argsItemSize), string(argsItemImage), string(argsAESKey), string(argsItemBasePrice), string(argsCurrentOwnerId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-item [id] [recType] [itemDesc] [itemDetail] [itemDate] [itemType] [itemSubject] [itemMedia] [itemSize] [itemImage] [AESKey] [itemBasePrice] [currentOwnerId]",
		Short: "Update a item",
		Args:  cobra.ExactArgs(13),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsItemDesc := string(args[2])
			argsItemDetail := string(args[3])
			argsItemDate := string(args[4])
			argsItemType := string(args[5])
			argsItemSubject := string(args[6])
			argsItemMedia := string(args[7])
			argsItemSize := string(args[8])
			argsItemImage := string(args[9])
			argsAESKey := string(args[10])
			argsItemBasePrice := string(args[11])
			argsCurrentOwnerId := string(args[12])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateItem(clientCtx.GetFromAddress().String(), id, string(argsRecType), string(argsItemDesc), string(argsItemDetail), string(argsItemDate), string(argsItemType), string(argsItemSubject), string(argsItemMedia), string(argsItemSize), string(argsItemImage), string(argsAESKey), string(argsItemBasePrice), string(argsCurrentOwnerId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-item [id] [recType] [itemDesc] [itemDetail] [itemDate] [itemType] [itemSubject] [itemMedia] [itemSize] [itemImage] [AESKey] [itemBasePrice] [currentOwnerId]",
		Short: "Delete a item by id",
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

			msg := types.NewMsgDeleteItem(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
