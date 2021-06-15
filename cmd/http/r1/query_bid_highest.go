package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctypes "github.com/jack139/artchain/x/auction/types"

	"github.com/cosmos/cosmos-sdk/client"

	"fmt"
	"strings"
	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)


/* 查询最高出价信息 */
func QueryBidHighest(ctx *fasthttp.RequestCtx) {
	log.Println("query_bid_highest")

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

	// 查询链上数据
	respData2, err := queryBidHighest(auctionIdStr)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	var resp map[string]interface{}

	if respData2==nil { // 无数据返回的情况
		resp = map[string] interface{} {
			"bid" : nil,
		}
	} else {
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

		resp = map[string] interface{} {
			"bid" : respData,
		}

	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryBidHighest(auctionIdStr string) (*map[string]interface{}, error) {
	// 获取 clientCtx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := auctypes.NewQueryClient(clientCtx)

	params := &auctypes.QueryGetHighBidRequest{
		AuctionId: auctionIdStr,
	}

	res, err := queryClient.BidHigh(context.Background(), params)
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

	if respData["Bid"]==nil { // 没有返回数据的情况
		return nil, nil
	}

	bidMap := respData["Bid"].(map[string]interface{})

	// 检查 lastDate 字段是否正常
	if _, ok := bidMap["lastDate"]; !ok {
		return nil, fmt.Errorf("lastDate empty") // 不应该发生
	}
	if !strings.HasPrefix(bidMap["lastDate"].(string), "[") {
		return nil, fmt.Errorf("lastDate broken") // 不应该发生
	}

	// 反序列化
	var data2 []map[string]interface{}
	if err := json.Unmarshal([]byte(bidMap["lastDate"].(string)), &data2); err != nil {
		return nil, err
	}
	bidMap["lastDate"] = data2

	return &(bidMap), nil
}
