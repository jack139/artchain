package r1

import (
	cmdclient "github.com/jack139/artchain/cmd/client"
	"github.com/jack139/artchain/cmd/http/helper"

	"github.com/valyala/fasthttp"
	"log"
)

/* 用户验证 */

func QueryUserVerify(ctx *fasthttp.RequestCtx) {
	log.Println("query_user_verify")

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
	mnemonic, ok := (*reqData)["mystery"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need mystery")
		return
	}

	// 验证用户
	verified, err := cmdclient.VerifyUserAccount(helper.HttpCmd, chainAddr, mnemonic)
	if err != nil {
		helper.RespError(ctx, 9009, err.Error())
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"verified": verified, // 是否验证通过
	}

	helper.RespJson(ctx, &resp)
}
