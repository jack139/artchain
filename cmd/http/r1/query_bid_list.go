package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctypes "github.com/jack139/artchain/x/auction/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询出价清单 */
func QueryBidList(ctx *fasthttp.RequestCtx) {
	log.Println("query_bid_list")

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
		helper.RespError(ctx, 9001, "need auction_id")
		return
	}
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9004, "need status")
		return
	}
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		helper.RespError(ctx, 9002, "need page")
		return
	}
	limit, ok := (*reqData)["limit"].(float64)
	if !ok {
		helper.RespError(ctx, 9003, "need limit")
		return
	}

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9005, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryBidListPage(ctx, auctionIdStr, uint64(page), uint64(limit), status)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		newItem := map[string]interface{} {
			"id"         : item["id"],
			"auction_id" : item["auctionId"],
			"bid_no"     : item["bidNo"],
			"buyer_addr" : item["buyerId"],
			"bid_price"  : item["bidPrice"],
			"bid_time"   : item["bidTime"],
			"last_date"  : item["lastDate"],
			"status"     : item["status"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"review_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryBidListPage(ctx *fasthttp.RequestCtx, auctionId string, page uint64, limit uint64, status string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := auctypes.NewQueryClient(clientCtx)

	pageReq := query.PageRequest{
		Key:        []byte(""),
		Offset:     (page - 1) * limit,
		Limit:      limit,
		CountTotal: true,
	}

	params := &auctypes.QueryAllBidRequest{
		AuctionId: auctionId,
		Status: status,
		Pagination: &pageReq,
	}

	res, err := queryClient.BidAll(context.Background(), params)
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

	itemMapList := respData["Bid"].([]interface{})

	return &(itemMapList), nil
}
