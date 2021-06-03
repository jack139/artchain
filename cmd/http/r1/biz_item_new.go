package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	invtypes "github.com/jack139/artchain/x/inventory/types"

	"github.com/gogo/protobuf/proto"

	"log"
	"bytes"
	"time"
	"strings"
	"strconv"
	"encoding/json"
	"encoding/hex"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
	//sdk "github.com/cosmos/cosmos-sdk/types"
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


	// 构建lastDate
	var lastDateMap []map[string]interface{}
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":  "new",
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
		"[]", //itemImage string, 
		"", //AESKey string, 
		itemBasePrice, //itemBasePrice string, 
		itemOwnerAddr, //currentOwnerId string, 
		"WAIT", //status string
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

	// 处理 data字段
	bs, err := hex.DecodeString(respData["data"].(string))
	if err != nil {
		helper.RespError(ctx, 9013, err.Error())
		return
	}

	// 0A100A0A4372656174654974656D1202081F
	// 0A   10   0A   0A  4372656174654974656D  12      02    081F
	// 标记 长度 标记  长度   数据                标记？  长度   数据
	slice1 := strings.Split(string(bs), "\n") // TODO: 有问题！ 需要解决格式分析问题
	//log.Println(len(slice1))
	//log.Println(slice1)
	slice2 := strings.Split(slice1[3], "\x12")
	//log.Println(len(slice2))
	//log.Println(slice2)
	//log.Println("new id: ", slice2[1])
	//log.Println("new id: %v", []byte(slice2[1])[1:])

	var pb invtypes.MsgCreateItemResponse
	//var pb sdk.Result
	if err := proto.Unmarshal([]byte(slice2[1])[1:], &pb); err != nil{
		helper.RespError(ctx, 9014, err.Error())
		return		
	}
	log.Println("New id: ", pb.Id)

	// 返回区块id
	resp := map[string]interface{}{
		"height" : respData["height"].(string),  // 区块高度
		"id" : strconv.FormatUint(pb.Id, 10), // item_id
	}

	helper.RespJson(ctx, &resp)
}
