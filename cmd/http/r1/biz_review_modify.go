package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	"context"
	"strconv"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

/* 修改物品评价信息 */

func BizReviewModify(ctx *fasthttp.RequestCtx) {
	log.Println("biz_review_modify")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	reviewIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
		return
	}
	reviewerAddr, ok := (*reqData)["reviewer_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need reviewer_addr")
		return
	}

	reviewDetail, _ := (*reqData)["detail"].(string)
	if len(reviewDetail)==0 { // 评论长度不能为0
		helper.RespError(ctx, 9003, "need detail")
		return		
	}

	reviewId, _ := strconv.ParseUint(reviewIdStr, 10, 64)

	// 获取当前链上数据
	reviewMap, err := queryReviewInfoById(ctx, reviewId)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	// 检查所有人addr是否一致
	if reviewerAddr!=(*reviewMap)["reviewerId"].(string) {
		helper.RespError(ctx, 9003, "wrong reviewer_addr")
		return				
	}

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 数据上链
	msg := invtypes.NewMsgUpdateReview(
		(*reviewMap)["creator"].(string), //creator string, 
		reviewId, //id uint64, 
		(*reviewMap)["recType"].(string), //recType string, 
		(*reviewMap)["itemId"].(string), //itemId string, 
		(*reviewMap)["reviewerId"].(string), //reviewerId string, 
		reviewDetail, //reviewDetail string, 
		(*reviewMap)["reviewDate"].(string), //reviewDate string, 
		(*reviewMap)["upCount"].(string), //upCount string, 
		(*reviewMap)["downCount"].(string), //downCount string,
	)
	if err := msg.ValidateBasic(); err != nil {
		helper.RespError(ctx, 9010, err.Error())
		return
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	err = tx.GenerateOrBroadcastTxCLI(clientCtx, helper.HttpCmd.Flags(), msg)
	if err != nil {
		helper.RespError(ctx, 9011, err.Error())
		return		
	}

	// 结果输出
	respBytes := []byte(buf.String())

	log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		helper.RespError(ctx, 9012, err.Error())
		return
	}

	// code==0 提交成功
	if respData["code"].(float64)!=0 { 
		helper.RespError(ctx, 9099, buf.String())  ///  提交失败
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"height" : respData["height"].(string),  // 区块高度
	}

	helper.RespJson(ctx, &resp)
}


// 查询链上数据, 返回 map
func queryReviewInfoById(ctx *fasthttp.RequestCtx, reviewId uint64) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetReviewRequest{
		Id: reviewId,
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

	log.Println("output: ", buf.String())

	//log.Printf("%v\n", string(respBytes))

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	itemMap := respData["Review"].(map[string]interface{})

	return &(itemMap), nil
}
