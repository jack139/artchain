package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询物品清单 */
func QueryItemList(ctx *fasthttp.RequestCtx) {
	log.Println("query_item_list")

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
	ownerAddr, ok := (*reqData)["owner_addr"].(string)

	if page < 1 || limit < 1 {
		helper.RespError(ctx, 9003, "page and limit need begin from 1")
		return		
	}

	// 查询链上数据
	respData2, err := queryItemListPage(ctx, uint64(page), uint64(limit), ownerAddr)
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
			"desc"       : item["itemDesc"],
			"detail"     : item["itemDetail"],
			"date"       : item["itemDate"],
			"type"       : item["itemType"],
			"subject"    : item["itemSubject"],
			"media"      : item["itemMedia"],
			"size"       : item["itemSize"],
			"base_price" : item["itemBasePrice"],
			"owner_addr" : item["currentOwnerId"],
			//"last_date"  : item["lastDate"],
			"status"     : item["status"],
		}
		respData = append(respData, newItem)
	}

	resp := map[string] interface{} {
		"item_list" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryItemListPage(ctx *fasthttp.RequestCtx, page uint64, limit uint64, ownerAddr string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	pageReq := query.PageRequest{
		Key:        []byte(""),
		Offset:     (page - 1) * limit,
		Limit:      limit,
		CountTotal: true,
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	if len(ownerAddr)==0 { // 查所有的
		params := &invtypes.QueryAllItemRequest{
			Pagination: &pageReq,
		}

		res, err := queryClient.ItemAll(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)

	} else { // 查指定owner_addr的
		params := &invtypes.QueryAllItemByOwnerRequest{
			CurrentOwnerId: ownerAddr,
			Pagination: &pageReq,
		}		

		res, err := queryClient.ItemAllByOwner(context.Background(), params)
		if err != nil {
			return nil, err
		}

		// 转换输出
		clientCtx.PrintProto(res)
	}

	//log.Printf("%T\n", res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMapList := respData["Item"].([]interface{})

	return &(itemMapList), nil
}
