package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	transtypes "github.com/jack139/artchain/x/trans/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询交易信息清单 */
// cate取值: seller, buyer, item, status
func QueryTransListByCondition(ctx *fasthttp.RequestCtx) {
	log.Println("query_trans_list_by_condition")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	cate, ok := (*reqData)["cate"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need cate")
		return
	}
	condition, ok := (*reqData)["condition"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need condition")
		return
	}
	page, ok := (*reqData)["page"].(float64)
	if !ok {
		helper.RespError(ctx, 9003, "need page")
		return
	}
	limit, ok := (*reqData)["limit"].(float64)
	if !ok {
		helper.RespError(ctx, 9004, "need limit")
		return
	}

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9005, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryTransListPageByCondition(ctx, cate, condition, uint64(page), uint64(limit))
	if err!=nil{
		helper.RespError(ctx, 9010, err.Error())
		return
	}	

	dataList := *respData2

	// 构建返回结构
	respData := make([]map[string]interface{}, 0) 

	for _, item0 := range dataList {
		item := item0.(map[string]interface{})

		newItem := map[string]interface{} {
			"id"           : item["id"],
			"auction_id"   : item["auctionId"],
			"item_id"      : item["itemId"],
			"trans_type"   : item["transType"],
			"buyer_addr"   : item["userId"],
			"trans_date"   : item["transDate"],
			"hammer_time"  : item["hammerTime"],
			"hammer_price" : item["hammerPrice"],
			"details"      : item["details"],
			"status"       : item["status"],
			//"last_date"    : item["lastDate"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"trans_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryTransListPageByCondition(ctx *fasthttp.RequestCtx, cate string, condition string,
	page uint64, limit uint64) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := transtypes.NewQueryClient(clientCtx)

	pageReq := query.PageRequest{
		Key:        []byte(""),
		Offset:     (page - 1) * limit,
		Limit:      limit,
		CountTotal: true,
	}

	params := &transtypes.QuerySomeTransactionRequest{
		Cate: cate,
		Condition: condition,
		Pagination: &pageReq,
	}

	res, err := queryClient.TransactionSome(context.Background(), params)
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

	transMapList := respData["Transaction"].([]interface{})

	return &(transMapList), nil
}
