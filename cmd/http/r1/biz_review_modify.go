package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	"context"
	"strconv"
	"bytes"
	"time"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
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
	callerAddr, ok := (*reqData)["caller_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9101, "need caller_addr")
		return
	}
	reviewIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
		return
	}
	itemIdStr, ok := (*reqData)["item_id"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need item_id")
		return
	}

	reviewDetail, _ := (*reqData)["detail"].(string)
	if len(reviewDetail)==0 { // 评论长度不能为0
		helper.RespError(ctx, 9004, "need detail")
		return		
	}

	reviewId, err := strconv.ParseUint(reviewIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	reviewMap, err := queryReviewInfoById(ctx, reviewId, itemIdStr)
	if err!=nil {
		helper.RespError(ctx, 9005, err.Error())
		return		
	}

	// 构建lastDate
	lastDateMap := (*reviewMap)["lastDate"].([]map[string]interface{})
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":  "edit",
		"date": time.Now().Format("2006-01-02 15:04:05"),
	})
	lastDate, err := json.Marshal(lastDateMap)
	if err != nil {
		helper.RespError(ctx, 9004, err.Error())
		return
	}

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9015, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*reviewMap)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

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
		"WAIT", // 设置为 wait
		string(lastDate), // lastDate
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
func queryReviewInfoById(ctx *fasthttp.RequestCtx, reviewId uint64, itemId string) (*map[string]interface{}, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetReviewRequest{
		Id: reviewId,
		ItemId: itemId,
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
