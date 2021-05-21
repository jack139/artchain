package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

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

/* 审核用户（修改用户状态） */

func BizAuditUser(ctx *fasthttp.RequestCtx) {
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
	callerAddr, ok := (*reqData)["caller_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9101, "need caller_addr")
		return
	}
	chainAddr, ok := (*reqData)["chain_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need chain_addr")
		return
	}
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9005, "need status")
		return
	}

	// 获取当前链上数据
	userMap, err := queryUserInfoByChainAddr(ctx, chainAddr)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}


	userInfoMap := (*userMap)["userInfo"].(map[string]interface{})
	userInfo, err := json.Marshal(userInfoMap)
	if err != nil {
		helper.RespError(ctx, 9003, err.Error())
		return
	}


	// 构建lastDate
	lastDateMap := (*userMap)["lastDate"].([]map[string]interface{})
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": callerAddr,
		"act":  "audit",
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
	helper.HttpCmd.Flags().Set(flags.FlagFrom, chainAddr)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	userId, err := strconv.ParseUint((*userMap)["id"].(string), 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 数据上链
	msg := persontypes.NewMsgUpdateUser(
		(*userMap)["creator"].(string), 
		userId, 
		(*userMap)["recType"].(string), 
		(*userMap)["name"].(string), 
		(*userMap)["userType"].(string), 
		string(userInfo), 
		status, // 其实 只修改了这里
		(*userMap)["regDate"].(string), 
		chainAddr,
		string(lastDate), // 和， 修改了这里
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
