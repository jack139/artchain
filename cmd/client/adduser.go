package client

import (
	"bufio"
	"log"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	bip39 "github.com/cosmos/go-bip39"

	"github.com/jack139/artchain/x/artchain/types"
)

// 建新用户 user， 建key，建account
// 返回： address, mnemonic
func AddUserAccount(cmd *cobra.Command, name string, reward string) (string, string, error) {
	// 保存 --from 设置
	originFlagFrom, err := cmd.Flags().GetString(flags.FlagFrom)
	if err != nil {
		return "", "", err
	}
	log.Println("FlagFrom:", originFlagFrom)

	// 设置 faucet 地址，用于转账
	cmd.Flags().Set(flags.FlagFrom, types.FaucetAddress) 
	// 结束时恢复 --from 设置
	defer cmd.Flags().Set(flags.FlagFrom, originFlagFrom) 

	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return "", "", err
	}

	// 获取 keyring 环境
	var kb keyring.Keyring

	buf := bufio.NewReader(cmd.InOrStdin())
	keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
	if err != nil {
		return "", "", err
	}
	kb, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.KeyringDir, buf)

	// 注册新的 key
	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return "", "", err
	}

	hdPath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()

	// read entropy seed straight from tmcrypto.Rand and convert to mnemonic
	mnemonicEntropySize := 256
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", "", err
	}

	// Get bip39 mnemonic
	var mnemonic, bip39Passphrase string

	mnemonic, err = bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", "", err
	}

	info, err := kb.NewAccount(name, mnemonic, bip39Passphrase, hdPath, algo)
	if err != nil {
		return "", "", err
	}

	log.Println("mnemonic: ", mnemonic)
	//log.Println(info)

	// 新用户的 地址
	toAddr := info.GetAddress()

	// 转账 1credit， 会自动建立auth的账户
	coins, err := sdk.ParseCoinsNormalized(reward)
	if err != nil {
		return "", "", err
	}

	msg := banktypes.NewMsgSend(clientCtx.GetFromAddress(), toAddr, coins)
	if err := msg.ValidateBasic(); err != nil {
		return "", "", err
	}

	// 参考cosmos-sdk/client/keys/show.go 中 getBechKeyOut()
	ko_new, err := keyring.Bech32KeyOutput(info)
	if err != nil {
		return "", "", err
	}

	// 取得地址字符串： 例如 artchain1zfqgxtujvpy92prtzgmzs3ygta9y2cl3w8hxlh
	addr_new := ko_new.Address

	// 调用 send 的 RPC 服务
	return addr_new, mnemonic, tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

}


// 比对用户 user， 
// 返回： bool
func VerifyUserAccount(cmd *cobra.Command, userAddr string, mnemonic string) (bool, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return false, err
	}

	// 获取 keyring 环境
	var kb keyring.Keyring

	buf := bufio.NewReader(cmd.InOrStdin())
	keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
	if err != nil {
		return false, err
	}
	kb, err = keyring.New(sdk.KeyringServiceName(), keyringBackend, clientCtx.KeyringDir, buf)

	// 注册新的 key
	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return false, err
	}

	hdPath := hd.CreateHDPath(sdk.GetConfig().GetCoinType(), 0, 0).String()

	// Get bip39 mnemonic
	var bip39Passphrase string

	if !bip39.IsMnemonicValid(mnemonic) && mnemonic != "" {
		return false, fmt.Errorf("invalid mnemonic")
	}

	// 生成私钥
	derivedPriv, err := algo.Derive()(mnemonic, bip39Passphrase, hdPath)
	if err != nil {
		return false, err
	}
	privKey := algo.Generate()(derivedPriv)

	// 从公钥生成acc地址
	accAddr := sdk.AccAddress(privKey.PubKey().Address().Bytes())
	//fmt.Println(accAddr.String())

	return accAddr.String()==userAddr, nil
}


/* 通过key name获取地址 */
func GetAddrStr(cmd *cobra.Command, keyref string) (string, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return "", err
	}

	// 获取 keyring 环境
	var kb keyring.Keyring

	buf := bufio.NewReader(cmd.InOrStdin())
	// keyringBackend 直接使用 test
	kb, err = keyring.New(sdk.KeyringServiceName(), "test", clientCtx.KeyringDir, buf)

	// 获取 地址
	//keyref := "faucet"
	info0, err := kb.Key(keyref)
	if err != nil {
		return "", err
	}
	//addr0 := info0.GetAddress() // AccAddress

	// 参考cosmos-sdk/client/keys/show.go 中 getBechKeyOut()
	ko, err := keyring.Bech32KeyOutput(info0)
	if err != nil {
		return "", err
	}

	// 取得地址字符串： 例如 artchain1zfqgxtujvpy92prtzgmzs3ygta9y2cl3w8hxlh
	addr0 := ko.Address
	//fmt.Println(addr0)

	return addr0, nil
}
