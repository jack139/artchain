package cli

import (
	"bufio"
	"errors"
	"fmt"
	//"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	//"github.com/spf13/viper"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/jack139/artchain/x/faucet/internal/types"
)

// GetTxCmd return faucet sub-command for tx
func GetTxCmd(cdc *codec.LegacyAmino) *cobra.Command {
	faucetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "faucet transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	//faucetTxCmd.AddCommand(flags.PostCommands(
	//	GetCmdMint(),
	//	GetCmdMintFor(),
	//	GetCmdInitial(),
	//	GetPublishKey(),
	//)...)

	faucetTxCmd.AddCommand(GetCmdMint())
	faucetTxCmd.AddCommand(GetCmdMintFor())
	faucetTxCmd.AddCommand(GetCmdInitial(cdc))
	faucetTxCmd.AddCommand(GetPublishKey())

	return faucetTxCmd
}

// GetCmdWithdraw is the CLI command for mining coin
func GetCmdMint() *cobra.Command {
	return &cobra.Command{
		Use:   "mint",
		Short: "mint coin to sender address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			/*
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgMint(cliCtx.GetFromAddress(), cliCtx.GetFromAddress(), time.Now().Unix())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			*/

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(clientCtx.GetFromAddress(), clientCtx.GetFromAddress(), time.Now().Unix())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdWithdraw is the CLI command for mining coin
func GetCmdMintFor() *cobra.Command {
	return &cobra.Command{
		Use:   "mintfor [address]",
		Short: "mint coin for new address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			/*
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			address, _ := sdk.AccAddressFromBech32(args[0])

			msg := types.NewMsgMint(cliCtx.GetFromAddress(), address, time.Now().Unix())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			*/

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			address, _ := sdk.AccAddressFromBech32(args[0])

			msg := types.NewMsgMint(clientCtx.GetFromAddress(), address, time.Now().Unix())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}
}

func GetPublishKey() *cobra.Command {
	return &cobra.Command{
		Use:   "publish",
		Short: "Publish current account as an public faucet. Do NOT add many coins in this account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			//inBuf := bufio.NewReader(cmd.InOrStdin())
			//cliCtx := context.NewCLIContext().WithCodec(cdc)

			//txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//kb, errkb := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), inBuf)
			//if errkb != nil {
			//	return errkb
			//}

			// 获取 keyring 环境
			var kb keyring.Keyring

			buf := bufio.NewReader(cmd.InOrStdin())
			keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
			if err != nil {
				return err
			}
			kb, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.KeyringDir, buf)

			// check local key
			armor, err := kb.ExportPubKeyArmor(clientCtx.GetFromName())
			if err != nil {
				return err
			}

			msg := types.NewMsgFaucetKey(clientCtx.GetFromAddress(), armor)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func GetCmdInitial(cdc *codec.LegacyAmino) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize mint key for faucet",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			//inBuf := bufio.NewReader(cmd.InOrStdin())
			//cliCtx := context.NewCLIContext().WithCodec(cdc)

			//txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//kb, errkb := keys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), inBuf)
			//if errkb != nil {
			//	return errkb
			//}

			// 获取 keyring 环境
			var kb keyring.Keyring

			buf := bufio.NewReader(cmd.InOrStdin())
			keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
			if err != nil {
				return err
			}
			kb, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.KeyringDir, buf)


			// check local key
			//_, err = kb.Get(types.ModuleName)
			_, err = kb.Key(types.ModuleName)
			if err == nil {
				return errors.New("faucet existed")
			}

			// fetch from chain
			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/key", types.ModuleName), nil)
			if err != nil {
				return nil
			}
			var rkey types.FaucetKey
			cdc.MustUnmarshalJSON(res, &rkey)

			if len(rkey.Armor) == 0 {
				return errors.New("Faucet key has not published")
			}
			// import to keybase
			kb.ImportPubKey(types.ModuleName, rkey.Armor)
			fmt.Println("The faucet has been loaded successfully.")
			return nil

		},
	}
}
