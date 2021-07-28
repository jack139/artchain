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

/* 查询评论信息 */
func QueryReviewInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_review_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	itemIdStr, ok := (*reqData)["item_id"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need item_id")
		return
	}
	reviewIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need id")
		return
	}

	reviewId, err := strconv.ParseUint(reviewIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 查询链上数据
	respData2, err := queryReviewInfoById(reviewId, itemIdStr)
	if err != nil {
		helper.RespError(ctx, 9014, err.Error())
		return
	}

	reviewMap := *respData2

	// 构建返回结构
	respData := map[string]interface{}{
		"id":            reviewMap["id"],
		"item_id":       reviewMap["itemId"],
		"detail":        reviewMap["reviewDetail"],
		"reviewer_addr": reviewMap["reviewerId"],
		"review_date":   reviewMap["reviewDate"],
		"last_date":     reviewMap["lastDate"],
		"status":        reviewMap["status"],
	}

	resp := map[string]interface{}{
		"review": respData,
	}

	helper.RespJson(ctx, &resp)
}

// 查询链上数据, 返回 map
func queryReviewInfoById(reviewId uint64, itemIdStr string) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetReviewRequest{
		Id:     reviewId,
		ItemId: itemIdStr,
	}

	res, err := queryClient.Review(context.Background(), params)
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

	reviewMap := respData["Review"].(map[string]interface{})

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
