package r1

import (
	cmdclient "github.com/jack139/artchain/cmd/client"
	"github.com/jack139/artchain/x/artchain/types"
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"log"
	"time"
	"bytes"
	"strings"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

/* 用户注册 */

func BizRegister(ctx *fasthttp.RequestCtx) {
	log.Println("biz_user_register")

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
	userName, ok := (*reqData)["login_name"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need login_name")
		return
	}
	userType, ok := (*reqData)["user_type"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need user_type")
		return
	}
	bank_acc_name, _ := (*reqData)["bank_acc_name"].(string)
	bank_name, _ := (*reqData)["bank_name"].(string)
	bank_acc_no, _ := (*reqData)["bank_acc_no"].(string)
	contact_address, _ := (*reqData)["address"].(string)
	phone, _ := (*reqData)["phone"].(string)
	email, _ := (*reqData)["email"].(string)
	referrer, _ := (*reqData)["referrer"].(string)

	// TODO： 检查 referrer 地址合法性

	// 构建userInfo
	userInfoMap := map[string]interface{}{
		"bank_acc_name": bank_acc_name,
		"bank_name":  bank_name,
		"bank_acc_no": bank_acc_no,
		"contact_address": contact_address,
		"phone": phone,
		"email": email,
		"referrer": referrer,
	}

	userInfo, err := json.Marshal(userInfoMap)
	if err != nil {
		helper.RespError(ctx, 9003, err.Error())
		return
	}

	// 初始化用户状态
	userStatus := "WAIT"
	if strings.HasPrefix(userType, "TRD") {
		userStatus = "ACTIVE"
	}

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


	// 生成新用户密钥
	address, mnemonic, err := cmdclient.AddUserAccount(helper.HttpCmd, userName, types.RewardRegister)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 设置 caller_addr
	originFlagFrom, err := helper.HttpCmd.Flags().GetString(flags.FlagFrom) // 保存 --from 设置
	if err != nil {
		helper.RespError(ctx, 9015, err.Error())
		return
	}
	helper.HttpCmd.Flags().Set(flags.FlagFrom, address)  // 设置 --from 地址
	defer helper.HttpCmd.Flags().Set(flags.FlagFrom, originFlagFrom)  // 结束时恢复 --from 设置

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 创建者地址，如果在生成新用户后，会变成faucet的地址
	//creatorAddr := clientCtx.GetFromAddress().String()

	// 数据上链
	msg := persontypes.NewMsgCreateUser(
		address, // creator string, 
		"USER", // recType string, 
		userName, // name string, 
		userType, // userType string, 
		string(userInfo), // userInfo string, 
		userStatus, // status string, 
		time.Now().Format("2006-01-02 15:04:05"), // regDate string, 
		address, // chainAddr string,
		string(lastDate), // lastDate string,
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
		"chain_addr" : address,  // 用户地址
		"mystery" : mnemonic, // 机密串
	}

	helper.RespJson(ctx, &resp)
}
