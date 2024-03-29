package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	transtypes "github.com/jack139/artchain/x/trans/types"

	"log"
	"bytes"
	"time"
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 新建成交交易 */
func BizTransNew(ctx *fasthttp.RequestCtx) {
	log.Println("biz_trans_new")

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
	buyerAddr, ok := (*reqData)["buyer_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need buyer_addr")
		return
	}
	auctionIdStr, ok := (*reqData)["auction_id"].(string)
	if !ok {
		helper.RespError(ctx, 9003, "need auction_id")
		return
	}
	itemIdStr, ok := (*reqData)["item_id"].(string)
	if !ok {
		helper.RespError(ctx, 9004, "need item_id")
		return
	}
	transType, ok := (*reqData)["trans_type"].(string)
	if !ok {
		helper.RespError(ctx, 9005, "need trans_type")
		return
	}
	hammerTime, ok := (*reqData)["hammer_time"].(string)
	if !ok {
		helper.RespError(ctx, 9006, "need hammer_time")
		return
	}
	hammerPrice, ok := (*reqData)["hammer_price"].(string)
	if !ok {
		helper.RespError(ctx, 9007, "need hammer_price")
		return
	}

	details, _ := (*reqData)["details"].(string)

	// TODO： 检查 buyerAddr 合法性, 

	auctionId, err := strconv.ParseUint(auctionIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9008, err.Error())
		return
	}

	// 获取当前链上数据, 拍卖信息, 进而获取卖家信息
	auctionMap, err := queryAuctionInfoById(auctionId)
	if err!=nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 检查拍卖状态是否是 WAIT， 其他状态不能修改
	if (*auctionMap)["status"].(string)!="CLOSE" {
		helper.RespError(ctx, 9010, "cannot create new transaction, status is not CLOSE")
		return
	}

	// 新建链上数据
	respData, err := transNew(callerAddr, auctionIdStr, itemIdStr, transType,
		buyerAddr, (*auctionMap)["SellerId"].(string), hammerTime, hammerPrice, details, "new")
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


func transNew(callerAddr string, auctionId string, itemId string, transType string, 
	buyerAddr string, sellerId string, hammerTime string, hammerPrice string, details string, 
	logText string) (*map[string]interface{}, error) {

	/* 信号量 */
	helper.AcquireSem(buyerAddr)
	defer helper.ReleaseSem(buyerAddr)

	// 构建lastDate
	var lastDateMap []map[string]interface{}
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":  logText,
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
	helper.HttpCmd.Flags().Set(flags.FlagFrom, buyerAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		return nil, err
	}

	// 数据上链
	msg := transtypes.NewMsgCreateTransaction(
		buyerAddr, //creator string, 
		"POSTTRAN", //recType string, 
		auctionId, //auctionId string, 
		itemId, //itemId string, 
		transType, //transType string, 
		buyerAddr, //buyerId string, 
		sellerId, //sellerId string, 
		time.Now().Format("2006-01-02 15:04:05"), //transDate string, 
		hammerTime, //hammerTime string, 
		hammerPrice, //hammerPrice string, 
		details, //details string, 
		"WAIT", //status string,
		string(lastDate),
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
