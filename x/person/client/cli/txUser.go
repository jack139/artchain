package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jack139/artchain/x/person/types"
)

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user [recType] [name] [userType] [address] [phone] [email] [bank] [accountNo] [status] [regDate]",
		Short: "Creates a new user",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsRecType := string(args[0])
			argsName := string(args[1])
			argsUserType := string(args[2])
			argsAddress := string(args[3])
			argsPhone := string(args[4])
			argsEmail := string(args[5])
			argsBank := string(args[6])
			argsAccountNo := string(args[7])
			argsStatus := string(args[8])
			argsRegDate := string(args[9])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(clientCtx.GetFromAddress().String(), string(argsRecType), string(argsName), string(argsUserType), string(argsAddress), string(argsPhone), string(argsEmail), string(argsBank), string(argsAccountNo), string(argsStatus), string(argsRegDate))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user [id] [recType] [name] [userType] [address] [phone] [email] [bank] [accountNo] [status] [regDate]",
		Short: "Update a user",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsRecType := string(args[1])
			argsName := string(args[2])
			argsUserType := string(args[3])
			argsAddress := string(args[4])
			argsPhone := string(args[5])
			argsEmail := string(args[6])
			argsBank := string(args[7])
			argsAccountNo := string(args[8])
			argsStatus := string(args[9])
			argsRegDate := string(args[10])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUser(clientCtx.GetFromAddress().String(), id, string(argsRecType), string(argsName), string(argsUserType), string(argsAddress), string(argsPhone), string(argsEmail), string(argsBank), string(argsAccountNo), string(argsStatus), string(argsRegDate))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user [id] [recType] [name] [userType] [address] [phone] [email] [bank] [accountNo] [status] [regDate]",
		Short: "Delete a user by id",
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

			msg := types.NewMsgDeleteUser(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
