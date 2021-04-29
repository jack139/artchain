package r1

import (
	cmdclient "github.com/jack139/artchain/cmd/client"
	"github.com/jack139/artchain/x/artchain/types"
	"github.com/jack139/artchain/cmd/http/helper"
	persontypes "github.com/jack139/artchain/x/person/types"

	"log"
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

/* 企业链业务处理 */

/* 用户注册
action == 13
*/

func BizRegister(ctx *fasthttp.RequestCtx) {
	log.Println("biz_register")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	userName, ok := (*reqData)["user_name"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need user_name")
		return
	}
	userType, ok := (*reqData)["user_type"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need user_type")
		return
	}
	//referrer, _ := (*reqData)["referrer"].(string)

	// TODO: 检查 userName是否已经存在！

	// 获取 ctx 上下文
	clientCtx, err := client.GetClientTxContext(helper.HttpCmd)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 创建者地址，如果在生成新用户后，会变成faucet的地址
	creatorAddr := clientCtx.GetFromAddress().String()


	// 生成新用户密钥
	address, mnemonic, err := cmdclient.AddUserAccount(helper.HttpCmd, userName, types.RewardRegister)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}


	// 数据上链
	msg := persontypes.NewMsgCreateUser(
		creatorAddr, // creator string, 
		"USER", // recType string, 
		userName, // name string, 
		userType, // userType string, 
		"", // address string, 
		"", // phone string, 
		"", // email string, 
		"", // bank string, 
		"", // accountNo string, 
		"ACTIVE", // status string, 
		"", // regDate string, 
		address, // chainAddr string,
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
		"address" : address,  // 用户地址
		"mystery" : mnemonic, // 机密串
	}

	helper.RespJson(ctx, &resp)
}
