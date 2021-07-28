package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctiontypes "github.com/jack139/artchain/x/auction/types"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"strings"
)

/* 查询拍卖信息 */
func QueryAuctionInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_auction_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	auctionIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need id")
		return
	}

	auctionId, err := strconv.ParseUint(auctionIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 查询链上数据
	respData2, err := queryAuctionInfoById(auctionId)
	if err != nil {
		helper.RespError(ctx, 9014, err.Error())
		return
	}

	auctionMap := *respData2

	// 构建返回结构
	respData := map[string]interface{}{
		"id":               auctionMap["id"],
		"item_id":          auctionMap["itemId"],
		"auction_house_id": auctionMap["auctionHouseId"],
		"seller_addr":      auctionMap["SellerId"],
		"req_date":         auctionMap["requestDate"],
		"reserved_price":   auctionMap["reservePrice"],
		"status":           auctionMap["status"],
		"open_date":        auctionMap["openDate"],
		"close_date":       auctionMap["closeDate"],
		"last_date":        auctionMap["lastDate"],
	}

	resp := map[string]interface{}{
		"auction": respData,
	}

	helper.RespJson(ctx, &resp)
}

// 查询链上数据, 返回 map
func queryAuctionInfoById(auctionId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := auctiontypes.NewQueryClient(clientCtx)

	params := &auctiontypes.QueryGetRequestRequest{
		Id: auctionId,
	}

	res, err := queryClient.Request(context.Background(), params)
	if err != nil {
		return nil, err
	}

	//log.Printf("%T\n", res)

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	// 转换输出
	clientCtx.PrintProto(res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	auctionMap := respData["Request"].(map[string]interface{})

	// 检查 lastDate 字段是否正常
	if _, ok := auctionMap["lastDate"]; !ok {
		return nil, fmt.Errorf("lastDate empty") // 不应该发生
	}
	if !strings.HasPrefix(auctionMap["lastDate"].(string), "[") {
		return nil, fmt.Errorf("lastDate broken") // 不应该发生
	}

	// 反序列化
	var data2 []map[string]interface{}
	if err := json.Unmarshal([]byte(auctionMap["lastDate"].(string)), &data2); err != nil {
		return nil, err
	}
	auctionMap["lastDate"] = data2

	return &(auctionMap), nil
}
