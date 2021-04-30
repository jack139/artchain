package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"github.com/cosmos/cosmos-sdk/client"

	"bytes"
	"context"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)



/* 查询物品信息 */
func QueryItemInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_item_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	itemIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need id")
		return
	}

	itemId, _ := strconv.ParseUint(itemIdStr, 10, 64)

	// 查询链上数据
	respData2, err := queryItemInfoById(ctx, itemId)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	itemMap := *respData2

	// 构建返回结构
	respData := map[string]interface{} {
		"desc"       : itemMap["itemDesc"],
		"detail"     : itemMap["itemDetail"],
		"date"       : itemMap["itemDate"],
		"type"       : itemMap["itemType"],
		"subject"    : itemMap["itemSubject"],
		"media"      : itemMap["itemMedia"],
		"size"       : itemMap["itemSize"],
		"base_price" : itemMap["itemBasePrice"],
		"owner_addr" : itemMap["currentOwnerId"],
		"last_date"  : itemMap["lastDate"],
	}

	resp := map[string] interface{} {
		"user" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryItemInfoById(ctx *fasthttp.RequestCtx, itemId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetItemRequest{
		Id: itemId,
	}

	res, err := queryClient.Item(context.Background(), params)
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

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMap := respData["Item"].(map[string]interface{})

	return &(itemMap), nil
}
