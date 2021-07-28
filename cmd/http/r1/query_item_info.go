package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

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

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 查询链上数据
	respData2, err := queryItemInfoById(itemId)
	if err != nil {
		helper.RespError(ctx, 9014, err.Error())
		return
	}

	itemMap := *respData2

	// 构建返回结构
	respData := map[string]interface{}{
		"id":         itemMap["id"],
		"desc":       itemMap["itemDesc"],
		"detail":     itemMap["itemDetail"],
		"date":       itemMap["itemDate"],
		"type":       itemMap["itemType"],
		"subject":    itemMap["itemSubject"],
		"media":      itemMap["itemMedia"],
		"size":       itemMap["itemSize"],
		"base_price": itemMap["itemBasePrice"],
		"owner_addr": itemMap["currentOwnerId"],
		"image":      itemMap["itemImage"],
		"last_date":  itemMap["lastDate"],
		"status":     itemMap["status"],
	}

	resp := map[string]interface{}{
		"item": respData,
	}

	helper.RespJson(ctx, &resp)
}

// 查询链上数据, 返回 map
func queryItemInfoById(itemId uint64) (*map[string]interface{}, error) {
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

	//log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMap := respData["Item"].(map[string]interface{})

	// 检查 image 字段是否正常
	if _, ok := itemMap["itemImage"]; !ok {
		return nil, fmt.Errorf("itemImage empty") // 不应该发生
	}
	if !strings.HasPrefix(itemMap["itemImage"].(string), "[") {
		itemMap["itemImage"] = "[]" // 不应该发生
	}

	// 反序列化
	var data1 []string
	if err := json.Unmarshal([]byte(itemMap["itemImage"].(string)), &data1); err != nil {
		return nil, err
	}
	itemMap["itemImage"] = data1

	// 检查 lastDate 字段是否正常
	if _, ok := itemMap["lastDate"]; !ok {
		return nil, fmt.Errorf("lastDate empty") // 不应该发生
	}
	if !strings.HasPrefix(itemMap["lastDate"].(string), "[") {
		return nil, fmt.Errorf("lastDate broken") // 不应该发生
	}

	// 反序列化
	var data2 []map[string]interface{}
	if err := json.Unmarshal([]byte(itemMap["lastDate"].(string)), &data2); err != nil {
		return nil, err
	}
	itemMap["lastDate"] = data2

	return &(itemMap), nil
}
