package r1

import (
	//cmdclient "github.com/jack139/artchain/cmd/client"
	//"github.com/jack139/artchain/x/artchain/types"
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"log"
	//"time"
	"strconv"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

/* 修改用户信息 */

func BizUserModify(ctx *fasthttp.RequestCtx) {
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
	chainAddr, ok := (*reqData)["chain_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need chain_addr")
		return
	}
	bank_acc_name, _ := (*reqData)["bank_acc_name"].(string)
	bank_name, _ := (*reqData)["bank_name"].(string)
	bank_acc_no, _ := (*reqData)["bank_acc_no"].(string)
	contact_address, _ := (*reqData)["address"].(string)
	phone, _ := (*reqData)["phone"].(string)
	email, _ := (*reqData)["email"].(string)

	// 获取当前链上数据
	userMap, err := queryUserInfoByChainAddr(ctx, chainAddr)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	userInfoOld := (*userMap)["userInfo"].(map[string]interface{})

	// 构建userInfo
	userInfoMap := map[string]interface{}{
		"bank_acc_name": userInfoOld["bank_acc_name"],
		"bank_name":  userInfoOld["bank_name"],
		"bank_acc_no": userInfoOld["bank_acc_no"],
		"contact_address": userInfoOld["contact_address"],
		"phone": userInfoOld["phone"],
		"email": userInfoOld["email"],
		"referrer": userInfoOld["referrer"],
	}

	// 是否要修改？
	if len(bank_acc_name)>0 {
		userInfoMap["bank_acc_name"] = bank_acc_name
	}
	if len(bank_name)>0 {
		userInfoMap["bank_name"] = bank_name
	}
	if len(bank_acc_no)>0 {
		userInfoMap["bank_acc_no"] = bank_acc_no
	}
	if len(contact_address)>0 {
		userInfoMap["contact_address"] = contact_address
	}
	if len(phone)>0 {
		userInfoMap["phone"] = phone
	}
	if len(email)>0 {
		userInfoMap["email"] = email
	}

	userInfo, err := json.Marshal(userInfoMap)
	if err != nil {
		helper.RespError(ctx, 9003, err.Error())
		return
	}


	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	userId, _ := strconv.ParseUint((*userMap)["id"].(string), 10, 64)

	// 数据上链
	msg := persontypes.NewMsgUpdateUser(
		(*userMap)["creator"].(string), 
		userId, 
		(*userMap)["recType"].(string), 
		(*userMap)["name"].(string), 
		(*userMap)["userType"].(string), 
		string(userInfo), 
		(*userMap)["status"].(string), 
		(*userMap)["regDate"].(string), 
		chainAddr,
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
