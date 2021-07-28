package cli

import (
	"github.com/spf13/cobra"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/faucet2/types"
)

func CmdCreateMining() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-Mining [Minter] [LastTime] [Total]",
		Short: "Creates a new Mining",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMinter := string(args[0])
			argsLastTime := string(args[1])
			argsTotal := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMining(clientCtx.GetFromAddress().String(), string(argsMinter), string(argsLastTime), string(argsTotal))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMining() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-Mining [id] [Minter] [LastTime] [Total]",
		Short: "Update a Mining",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsMinter := string(args[1])
			argsLastTime := string(args[2])
			argsTotal := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMining(clientCtx.GetFromAddress().String(), id, string(argsMinter), string(argsLastTime), string(argsTotal))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMining() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-Mining [id] [Minter] [LastTime] [Total]",
		Short: "Delete a Mining by id",
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

			msg := types.NewMsgDeleteMining(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint",
		Short: "Mint coin to sender address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(clientCtx.GetFromAddress().String(),
				clientCtx.GetFromAddress().String(),
				uint64(time.Now().Unix()))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
