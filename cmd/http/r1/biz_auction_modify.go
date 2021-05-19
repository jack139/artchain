package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctiontypes "github.com/jack139/artchain/x/auction/types"

	"log"
	"strconv"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 修改拍卖 */
func BizAuctionModify(ctx *fasthttp.RequestCtx) {
	log.Println("biz_auction_modify")

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
	auctionIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
		return
	}
	auctionHouseId, ok := (*reqData)["auction_house_id"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need auction_house_id")
		return
	}
	reservedPrice, ok := (*reqData)["reserved_price"].(string)
	if !ok {
		helper.RespError(ctx, 9003, "need reserved_price")
		return
	}

	auctionId, err := strconv.ParseUint(auctionIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	auctionMap, err := queryAuctionInfoById(ctx, auctionId)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	// 检查拍卖状态是否是 WAIT， 其他状态不能修改
	if (*auctionMap)["status"].(string)!="WAIT" {
		helper.RespError(ctx, 9003, "cannot modify, status is not WAIT")
		return				
	}

	// 是否要修改？
	if len(auctionHouseId)==0 {
		auctionHouseId = (*auctionMap)["auctionHouseId"].(string)
	}
	if len(reservedPrice)==0 {
		reservedPrice = (*auctionMap)["reservePrice"].(string)
	}

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9015, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, callerAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 数据上链
	msg := auctiontypes.NewMsgUpdateRequest(
		(*auctionMap)["creator"].(string), //creator string, 
		auctionId, //id uint64, 
		(*auctionMap)["recType"].(string), //recType string, 
		(*auctionMap)["itemId"].(string), //itemId string, 
		auctionHouseId, //auctionHouseId string, 
		(*auctionMap)["SellerId"].(string), //SellerId string, 
		(*auctionMap)["requestDate"].(string), //requestDate string, 
		reservedPrice, //reservePrice string, 
		(*auctionMap)["status"].(string), //status string, 
		(*auctionMap)["openDate"].(string), //openDate string, 
		(*auctionMap)["closeDate"].(string), //closeDate string,
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
