package cmd

import (
	"log"
	"errors"
	"time"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	httpcli "github.com/jack139/artchain/cmd/http"
	httphelper "github.com/jack139/artchain/cmd/http/helper"
)

func HttpCliCmd() *cobra.Command {
	cmd := &cobra.Command{ // 启动http服务
		Use:   "http <port>",
		Short: "start http service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("need port number")
			}

			// 保存 cmd
			httphelper.HttpCmd = cmd

			// 设置定时任务
			ticker1 := time.NewTicker(5 * time.Second)
			// 一定要调用Stop()，回收资源
			defer ticker1.Stop()
			go func(t *time.Ticker) {
				for {
					// 每5秒中从chan t.C 中读取一次
					<-t.C
					// 检查拍卖时间，进行状态转换
					if err := checkAuction(); err!=nil{
						log.Println(err.Error())
					}
				}
			}(ticker1)

			// 启动 http 服务
			httpcli.RunServer(args[0])
			// 不会返回
			return nil
		},
	}

	cmd.Flags().String(flags.FlagChainID, "", "network chain ID")
	cmd.Flags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test|memory)")
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to tendermint rpc interface for this chain")
	cmd.Flags().BoolP(flags.FlagSkipConfirmation, "y", true, "Skip tx broadcasting prompt confirmation")

	return cmd
}
