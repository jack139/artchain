package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/legacy"

	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

/* 获取区块数据 */
func getBlock(clientCtx client.Context, height *int64) ([]byte, error) {
	// get the node
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	// header -> BlockchainInfo
	// header, tx -> Block
	// results -> BlockResults
	res, err := node.Block(context.Background(), height)
	if err != nil {
		return nil, err
	}

	return legacy.Cdc.MarshalJSON(res)
}

/* 指定区块查询交易 */
func QueryRawBlock(ctx *fasthttp.RequestCtx) {
	log.Println("query_block_rawdata")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	pubkey, ok := (*reqData)["chain_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need chain_addr")
		return
	}
	height, ok := (*reqData)["height"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need height")
		return
	}

	height64, err := strconv.ParseInt(height, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
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

	// 准备查询
	respBytes, err := getBlock(clientCtx, &height64)
	if err != nil {
		helper.RespError(ctx, 9006, err.Error())
		return
	}

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
