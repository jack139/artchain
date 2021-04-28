package r1

import (
	cmdclient "github.com/jack139/artchain/cmd/client"
	"github.com/jack139/artchain/x/artchain/types"
	"github.com/jack139/artchain/cmd/http/helper"

	"log"
	"github.com/valyala/fasthttp"
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
	//userType, ok := (*reqData)["user_type"].(string)
	//if !ok {
	//	respError(ctx, 9002, "need user_type")
	//	return
	//}
	//referrer, _ := (*reqData)["referrer"].(string)

	// 生成新用户密钥
	address, mnemonic, err := cmdclient.AddUserAccount(helper.HttpCmd, userName, types.RewardRegister)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"block":    map[string]interface{}{"id": ""}, // 为了兼容旧接口，目前无数据返回
		"userkey":  address,
		"mnemonic": mnemonic,
	}

	helper.RespJson(ctx, &resp)
}
