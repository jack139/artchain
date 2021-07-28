package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctypes "github.com/jack139/artchain/x/auction/types"

	"bytes"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"time"
)

/* 撤销出价 */

func BizAuctionBidWithdraw(ctx *fasthttp.RequestCtx) {
	log.Println("biz_auction_bid_withdraw")

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
	bidIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
		return
	}

	auctionIdStr, ok := (*reqData)["auction_id"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need auction_id")
		return
	}

	bidId, err := strconv.ParseUint(bidIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	bidMap, err := queryBidInfoById(bidId, auctionIdStr)
	if err != nil {
		helper.RespError(ctx, 9005, err.Error())
		return
	}

	/* 信号量 */
	helper.AcquireSem((*bidMap)["creator"].(string))
	defer helper.ReleaseSem((*bidMap)["creator"].(string))

	// 构建lastDate
	lastDateMap := (*bidMap)["lastDate"].([]map[string]interface{})
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":    "withdraw",
		"date":   time.Now().Format("2006-01-02 15:04:05"),
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
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*bidMap)["creator"].(string)) // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)          // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 数据上链
	msg := auctypes.NewMsgUpdateBid(
		(*bidMap)["creator"].(string),   //creator string,
		bidId,                           //id uint64,
		(*bidMap)["recType"].(string),   //recType string,
		(*bidMap)["auctionId"].(string), //auctionId string,
		(*bidMap)["bidNo"].(string),     //bidNo string,
		"WITHDRAW",                      //status string,
		(*bidMap)["buyerId"].(string),   //buyerId string,
		(*bidMap)["bidPrice"].(string),  //bidPrice string,
		(*bidMap)["bidTime"].(string),   //bidTime string,
		string(lastDate),                //lastDate string
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
	if respData["code"].(float64) != 0 {
		helper.RespError(ctx, 9099, buf.String()) ///  提交失败
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"height": respData["height"].(string), // 区块高度
	}

	helper.RespJson(ctx, &resp)
}
