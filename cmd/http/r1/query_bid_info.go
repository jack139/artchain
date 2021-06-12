package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctypes "github.com/jack139/artchain/x/auction/types"

	"github.com/cosmos/cosmos-sdk/client"

	"fmt"
	"strings"
	"bytes"
	"context"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)


/* 查询出价信息 */
func QueryBidInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_bid_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	auctionIdStr, ok := (*reqData)["auction_id"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need auction_id")
		return
	}
	bidIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need id")
		return
	}

	bidId, err := strconv.ParseUint(bidIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 查询链上数据
	respData2, err := queryBidInfoById(ctx, bidId, auctionIdStr)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	reviewMap := *respData2

	// 构建返回结构
	respData := map[string]interface{} {
		"id"         : reviewMap["id"],
		"auction_id" : reviewMap["auctionId"],
		"bid_no"     : reviewMap["bidNo"],
		"buyer_addr" : reviewMap["buyerId"],
		"bid_price"  : reviewMap["bidPrice"],
		"bid_time"   : reviewMap["bidTime"],
		"last_date"  : reviewMap["lastDate"],
		"status"     : reviewMap["status"],
	}

	resp := map[string] interface{} {
		"bid" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryBidInfoById(ctx *fasthttp.RequestCtx, bidId uint64, auctionIdStr string) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := auctypes.NewQueryClient(clientCtx)

	params := &auctypes.QueryGetBidRequest{
		Id: bidId,
		AuctionId: auctionIdStr,
	}

	res, err := queryClient.Bid(context.Background(), params)
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

	reviewMap := respData["Bid"].(map[string]interface{})

	// 检查 lastDate 字段是否正常
	if _, ok := reviewMap["lastDate"]; !ok {
		return nil, fmt.Errorf("lastDate empty") // 不应该发生
	}
	if !strings.HasPrefix(reviewMap["lastDate"].(string), "[") {
		return nil, fmt.Errorf("lastDate broken") // 不应该发生
	}

	// 反序列化
	var data2 []map[string]interface{}
	if err := json.Unmarshal([]byte(reviewMap["lastDate"].(string)), &data2); err != nil {
		return nil, err
	}
	reviewMap["lastDate"] = data2

	return &(reviewMap), nil
}
