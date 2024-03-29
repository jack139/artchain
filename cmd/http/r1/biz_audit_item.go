package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"log"
	"strconv"
	"github.com/valyala/fasthttp"
)

/* 审核物品（修改状态） */

func BizAuditItem(ctx *fasthttp.RequestCtx) {
	log.Println("biz_audit_item")

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
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need status")
		return
	}

	logText, ok := (*reqData)["action"].(string)
	if !ok {
		logText = "audit"
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 修改链上数据
	respData, err := itemModify(callerAddr, 
		itemId, "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", 
		"\x00", "\x00", "\x00", "\x00", status, logText)
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
