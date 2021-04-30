package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	//"time"
	"strconv"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

/* 修改物品信息 */

func BizItemModify(ctx *fasthttp.RequestCtx) {
	log.Println("biz_user_modify")

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
		helper.RespError(ctx, 9001, "need id")
		return
	}
	itemOwnerAddr, ok := (*reqData)["owner_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need owner_addr")
		return
	}

	itemDesc, ok := (*reqData)["desc"].(string)
	itemDate, ok := (*reqData)["date"].(string)
	itemDetail, _ := (*reqData)["detail"].(string)
	itemType, _ := (*reqData)["type"].(string)
	itemSubject, _ := (*reqData)["subject"].(string)
	itemMedia, _ := (*reqData)["media"].(string)
	itemSize, _ := (*reqData)["size"].(string)
	itemBasePrice, _ := (*reqData)["base_price"].(string)

	itemId, _ := strconv.ParseUint(itemIdStr, 10, 64)

	// 获取当前链上数据
	itemMap, err := queryItemInfoById(ctx, itemId)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	// 检查所有人addr是否一致
	if itemOwnerAddr!=(*itemMap)["currentOwnerId"].(string) {
		helper.RespError(ctx, 9003, "wrong owner_addr")
		return				
	}

	// 是否要修改？
	if len(itemDesc)==0 {
		itemDesc = (*itemMap)["itemDesc"].(string)
	}
	if len(itemDate)==0 {
		itemDate = (*itemMap)["itemDate"].(string)
	}
	if len(itemDetail)==0 {
		itemDetail = (*itemMap)["itemDetail"].(string)
	}
	if len(itemType)==0 {
		itemType = (*itemMap)["itemType"].(string)
	}
	if len(itemSubject)==0 {
		itemSubject = (*itemMap)["itemSubject"].(string)
	}
	if len(itemMedia)==0 {
		itemMedia = (*itemMap)["itemMedia"].(string)
	}
	if len(itemSize)==0 {
		itemSize = (*itemMap)["itemSize"].(string)
	}
	if len(itemBasePrice)==0 {
		itemBasePrice = (*itemMap)["itemBasePrice"].(string)
	}


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 数据上链
	msg := invtypes.NewMsgUpdateItem(
		(*itemMap)["creator"].(string), //creator string, 
		itemId, //id uint64, 
		(*itemMap)["recType"].(string), //recType string, 
		itemDesc, //itemDesc string, 
		itemDetail, //itemDetail string, 
		itemDate, //itemDate string, 
		itemType, //itemType string, 
		itemSubject, //itemSubject string, 
		itemMedia, //itemMedia string, 
		itemSize, //itemSize string, 
		(*itemMap)["itemImage"].(string), //itemImage string, 
		(*itemMap)["AESKey"].(string), //AESKey string, 
		itemBasePrice, //itemBasePrice string, 
		(*itemMap)["currentOwnerId"].(string), //currentOwnerId string, 
		(*itemMap)["status"].(string), //status string
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
