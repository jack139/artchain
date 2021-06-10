package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	"strconv"
	"bytes"
	"time"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 删除物品图片 - 并为实际从ipfs删除，只是从链上数据中删除 */

func IpfsRemoveImage(ctx *fasthttp.RequestCtx) {
	log.Println("ipfs_remove_image")

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
	itemIdStr, ok := (*reqData)["item_id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need item_id")
		return
	}

	hash, ok := (*reqData)["hash"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need hash")
		return
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	itemMap, err := queryItemInfoById(ctx, itemId)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	// 准备数据
	loadData := (*itemMap)["itemImage"].([]string)	
	//log.Printf("1: %v", loadData)
	for i, h := range loadData{
		if h==hash { // 删除 图片的  hash
			loadData = append(loadData[:i], loadData[i+1:]...)
			break
		}
	}
	//log.Printf("2: %v", loadData)
	loadBytes, err := json.Marshal(loadData)
	if err != nil {
		helper.RespError(ctx, 9008, err.Error())
		return
	}

	// 构建lastDate
	lastDateMap := (*itemMap)["lastDate"].([]map[string]interface{})
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":  "upload image",
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
	helper.HttpCmd.Flags().Set(flags.FlagFrom, (*itemMap)["creator"].(string))  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

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
		(*itemMap)["itemDesc"].(string), //itemDesc string, 
		(*itemMap)["itemDetail"].(string), //itemDetail string, 
		(*itemMap)["itemDate"].(string), //itemDate string, 
		(*itemMap)["itemType"].(string), //itemType string, 
		(*itemMap)["itemSubject"].(string), //itemSubject string, 
		(*itemMap)["itemMedia"].(string), //itemMedia string, 
		(*itemMap)["itemSize"].(string), //itemSize string, 
		string(loadBytes), //itemImage string, 图片信息， ipfs 哈希列表
		(*itemMap)["AESKey"].(string), //AESKey string, 
		(*itemMap)["itemBasePrice"].(string), //itemBasePrice string, 
		(*itemMap)["currentOwnerId"].(string), //currentOwnerId string, 
		"WAIT", // 修改后状态自动设置为 WAIT
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