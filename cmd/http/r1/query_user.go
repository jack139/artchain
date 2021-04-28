package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询用户余额 */
func QueryBalance(ctx *fasthttp.RequestCtx) {
	log.Println("query_raw_block")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	pubkey, ok := (*reqData)["userkey"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need userkey")
		return
	}

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9005, err.Error())
		return
	}

	// 检查 用户地址 是否存在
	_, err = helper.FetchKey(clientCtx.Keyring, pubkey)
	if err != nil {
		helper.RespError(ctx, 9001, "invalid userkey")
		return
	}

	addr, err := sdk.AccAddressFromBech32(pubkey)
	if err != nil {
		helper.RespError(ctx, 9006, err.Error())
		return
	}

	queryClient := banktypes.NewQueryClient(clientCtx)

	// 准备查询
	params := banktypes.NewQueryBalanceRequest(addr, "credit")
	res, err := queryClient.Balance(helper.HttpCmd.Context(), params)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	// 转换输出
	clientCtx.PrintProto(res.Balance)

	// 输出的字节流
	respBytes := []byte(buf.String())

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}
	if len(respBytes) > 0 {
		if err := json.Unmarshal(respBytes, &respData); err != nil {
			helper.RespError(ctx, 9004, err.Error())
			return
		}
	}
	resp := map[string]interface{}{
		"blcok": respData,
	}

	helper.RespJson(ctx, &resp)
}
