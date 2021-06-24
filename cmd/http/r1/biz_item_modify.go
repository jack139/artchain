package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"log"
	"strconv"
	"bytes"
	"fmt"
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 修改物品信息 */

func BizItemModify(ctx *fasthttp.RequestCtx) {
	log.Println("biz_item_modify")

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
	itemIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
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

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 修改链上数据
	respData, err := itemModify(callerAddr, 
		itemId, itemDesc, itemDetail, itemDate, itemType, 
		itemSubject, itemMedia, itemSize, "\x00", "\x00", 
		itemBasePrice, "\x00", "WAIT", "edit")
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

// string 参数填 "\x00" 表示不修改
func itemModify(callerAddr string, 
		itemId uint64, itemDesc string, itemDetail string, itemDate string, itemType string, 
		itemSubject string, itemMedia string, itemSize string, itemImage string, AESKey string, 
		itemBasePrice string, currentOwnerId string, status string, 
		logText string ) (*map[string]interface{}, error) {

	// 获取 creator
	creator, err := queryItemCreatorById(itemId)
	if err!=nil {
		return nil, err
	}

	/* 信号量 */
	helper.AcquireSem(creator)
	defer helper.ReleaseSem(creator)

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		return nil, err
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, creator)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		return nil, err
	}

	// 数据上链
	msg := invtypes.NewMsgUpdateItem(
		creator, //creator string, 
		itemId, //id uint64, 
		"\x00", //recType string, 
		itemDesc, //itemDesc string, 
		itemDetail, //itemDetail string, 
		itemDate, //itemDate string, 
		itemType, //itemType string, 
		itemSubject, //itemSubject string, 
		itemMedia, //itemMedia string, 
		itemSize, //itemSize string, 
		itemImage, //itemImage string, 
		AESKey, //AESKey string, 
		itemBasePrice, //itemBasePrice string, 
		currentOwnerId, //currentOwnerId string, 
		status, // status 修改后状态自动设置为 WAIT
		callerAddr+"|"+logText, // lastDate
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


// 查询链上数据, 返回 creator
func queryItemCreatorById(itemId uint64) (string, error) {
	// 获取 ctx 上下文
	clientCtx := client.GetClientContextFromCmd(helper.HttpCmd)

	// 准备查询
	queryClient := invtypes.NewQueryClient(clientCtx)

	params := &invtypes.QueryGetItemCreatorRequest{
		Id: itemId,
	}

	res, err := queryClient.ItemCreator(context.Background(), params)
	if err != nil {
		return "", err
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
		return "", err
	}

	return respData["creator"].(string), nil
}
