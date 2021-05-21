package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询物品清单 by status */
func QueryAuditItemList(ctx *fasthttp.RequestCtx) {
	log.Println("query_audit_item_list")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need status")
		return
	}

	// 查询链上数据
	respData2, err := queryAuditItemListPage(ctx, status)
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
func queryAuditItemListPage(ctx *fasthttp.RequestCtx, status string) (*[]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf


	params := &invtypes.QueryGetItemByStatusRequest{
		Status: status,
	}

	res, err := queryClient.ItemByStatus(context.Background(), params)
	if err != nil {
		return nil, err
	}

	// 转换输出
	clientCtx.PrintProto(res)

	//log.Printf("%T\n", res)

	// 输出的字节流
	respBytes := []byte(buf.String())

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMapList := respData["Item"].([]interface{})

	return &(itemMapList), nil
}
