package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctiontypes "github.com/jack139/artchain/x/auction/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询拍卖信息清单 */
func QueryAuctionList(ctx *fasthttp.RequestCtx) {
	log.Println("query_auction_list")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		helper.RespError(ctx, 9001, "need page")
		return
	}
	limit, ok := (*reqData)["limit"].(float64)
	if !ok {
		helper.RespError(ctx, 9002, "need limit")
		return
	}
	sellerAddr, ok := (*reqData)["seller_addr"].(string)

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9003, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryAuctionListPage(ctx, uint64(page), uint64(limit), sellerAddr)
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
			"id"               : item["id"],
			"item_id"          : item["itemId"],
			"auction_house_id" : item["auctionHouseId"],
			"seller_addr"      : item["SellerId"],
			"req_date"         : item["requestDate"],
			"reserved_price"   : item["reservePrice"],
			"status"           : item["status"],
			"open_date"        : item["openDate"],
			"close_date"       : item["closeDate"],
			//"last_date"        : item["lastDate"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"auction_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryAuctionListPage(ctx *fasthttp.RequestCtx, page uint64, limit uint64, sellerAddr string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := auctiontypes.NewQueryClient(clientCtx)

	pageReq := query.PageRequest{
		Key:        []byte(""),
		Offset:     (page - 1) * limit,
		Limit:      limit,
		CountTotal: true,
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	if len(sellerAddr)==0 { // 查所有的
		params := &auctiontypes.QueryAllRequestRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.RequestAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	} else { // 查指定seller_addr的
		params := &auctiontypes.QueryGetRequestByChainAddrRequest{
			ChainAddr  : sellerAddr,
			Pagination : &pageReq,
		}

		res, err := queryClient.RequestByChainAddr(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	}

	// 输出的字节流
	respBytes := []byte(buf.String())

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	auctionMapList := respData["Request"].([]interface{})

	return &(auctionMapList), nil
}
