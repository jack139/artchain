package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	auctiontypes "github.com/jack139/artchain/x/auction/types"

	"log"
	"strconv"
	"bytes"
	"time"
	"fmt"
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

	// 修改链上数据
	respData, err := auctionModify(auctionMap, callerAddr, auctionId, 
		auctionHouseId, reservedPrice, "", "", "", "edit")
	if err != nil {
		helper.RespError(ctx, 9010, err.Error())
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"height" : (*respData)["height"].(string),  // 区块高度
	}

	helper.RespJson(ctx, &resp)
}


func auctionModify(auctionMap *map[string]interface{}, callerAddr string, 
	auctionId uint64, auctionHouseId string, reservedPrice string, 
	openDate string, closeDate string, status string, logText string) (*map[string]interface{}, error) {

	// 为空串用原有值填充
	if len(auctionHouseId)==0 {
		auctionHouseId = (*auctionMap)["auctionHouseId"].(string)
	}
	if len(reservedPrice)==0 {
		reservedPrice = (*auctionMap)["reservePrice"].(string)
	}
	if len(openDate)==0 {
		openDate = (*auctionMap)["openDate"].(string)
	}
	if len(closeDate)==0 {
		closeDate = (*auctionMap)["closeDate"].(string)
	}
	if len(status)==0 {
		status = (*auctionMap)["status"].(string)
	}

	/* 信号量 */
	helper.AcquireSem((*auctionMap)["creator"].(string))
	defer helper.ReleaseSem((*auctionMap)["creator"].(string))

	// 构建lastDate
	lastDateMap := (*auctionMap)["lastDate"].([]map[string]interface{})
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act": logText,
		"date": time.Now().Format("2006-01-02 15:04:05"),
	})
	lastDate, err := json.Marshal(lastDateMap)
	if err != nil {
		return nil, err
	}

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		return nil, err
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*auctionMap)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		return nil, err
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
		status, //status string, 
		openDate, //openDate string, 
		closeDate, //closeDate string,
		string(lastDate), // lastDate
	)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// 设置 接收输出
	buf := new(bytes.Buffer)
	clientCtx.Output = buf

	err = tx.GenerateOrBroadcastTxCLI(clientCtx, helper.HttpCmd.Flags(), msg)
	if err != nil {
		return nil, err
	}

	// 结果输出
	respBytes := []byte(buf.String())

	log.Println("output: ", buf.String())

	// 转换成map, 生成返回数据
	var respData map[string]interface{}

	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return nil, err
	}

	// code==0 提交成功
	if respData["code"].(float64)!=0 { 
		return nil, fmt.Errorf("Tx fail: %s", buf.String())
	}

	return &respData, nil
}