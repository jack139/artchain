package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	transtypes "github.com/jack139/artchain/x/trans/types"

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


/* 查询交易信息 */
func QueryTransInfo(ctx *fasthttp.RequestCtx) {
	log.Println("query_trans_info")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	transIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9009, "need id")
		return
	}

	transId, err := strconv.ParseUint(transIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 查询链上数据
	respData2, err := queryTransInfoById(transId)
	if err!=nil{
		helper.RespError(ctx, 9014, err.Error())
		return
	}	

	transMap := *respData2

	// 构建返回结构
	respData := map[string]interface{} {
		"id"           : transMap["id"],
		"auction_id"   : transMap["auctionId"],
		"item_id"      : transMap["itemId"],
		"trans_type"   : transMap["transType"],
		"buyer_addr"   : transMap["userId"],
		"trans_date"   : transMap["transDate"],
		"hammer_time"  : transMap["hammerTime"],
		"hammer_price" : transMap["hammerPrice"],
		"details"      : transMap["details"],
		"status"       : transMap["status"],
		"last_date"    : transMap["lastDate"],
	}

	resp := map[string] interface{} {
		"trans" : respData,
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryTransInfoById(transId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := transtypes.NewQueryClient(clientCtx)

	params := &transtypes.QueryGetTransactionRequest{
		Id: transId,
	}

	res, err := queryClient.Transaction(context.Background(), params)
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

	transMap := respData["Transaction"].(map[string]interface{})


	// 检查 lastDate 字段是否正常
	if _, ok := transMap["lastDate"]; !ok {
		return nil, fmt.Errorf("lastDate empty") // 不应该发生
	}
	if !strings.HasPrefix(transMap["lastDate"].(string), "[") {
		return nil, fmt.Errorf("lastDate broken") // 不应该发生
	}

	// 反序列化
	var data2 []map[string]interface{}
	if err := json.Unmarshal([]byte(transMap["lastDate"].(string)), &data2); err != nil {
		return nil, err
	}
	transMap["lastDate"] = data2


	return &(transMap), nil
}
