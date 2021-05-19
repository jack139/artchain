package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 新建物品 */

func BizItemNew(ctx *fasthttp.RequestCtx) {
	log.Println("biz_item_new")

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
	itemDesc, ok := (*reqData)["desc"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need desc")
		return
	}
	itemDate, ok := (*reqData)["date"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need date")
		return
	}
	itemDetail, _ := (*reqData)["detail"].(string)
	itemType, _ := (*reqData)["type"].(string)
	itemSubject, _ := (*reqData)["subject"].(string)
	itemMedia, _ := (*reqData)["media"].(string)
	itemSize, _ := (*reqData)["size"].(string)
	itemBasePrice, _ := (*reqData)["base_price"].(string)
	itemOwnerAddr, _ := (*reqData)["owner_addr"].(string)

	// TODO： 检查 itemOwnerAddr 合法性


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

	// 创建者地址
	//creatorAddr := clientCtx.GetFromAddress().String()


	// 数据上链
	msg := invtypes.NewMsgCreateItem(
		callerAddr, //creator string, 
		"ARTINV", //recType string, 
		itemDesc, //itemDesc string, 
		itemDetail, //itemDetail string, 
		itemDate, //itemDate string, 
		itemType, //itemType string, 
		itemSubject, //itemSubject string, 
		itemMedia, //itemMedia string, 
		itemSize, //itemSize string, 
		"", //itemImage string, 
		"", //AESKey string, 
		itemBasePrice, //itemBasePrice string, 
		itemOwnerAddr, //currentOwnerId string, 
		"WAIT", //status string
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
