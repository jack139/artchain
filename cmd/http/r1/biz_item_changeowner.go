package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"log"
	"strconv"
	"github.com/valyala/fasthttp"
)

/* 修改物品所有人 */

func BizItemChangeOwner(ctx *fasthttp.RequestCtx) {
	log.Println("biz_item_change_owner")

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

	ownerAddr, ok := (*reqData)["owner_addr"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need owner_addr")
		return
	}

	// TODO：  检查 ownerAddr 合法性

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 修改链上数据
	respData, err := itemModify(callerAddr, 
		itemId, "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", 
		"\x00", "\x00", "\x00", ownerAddr, "\x00", "change owner")
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
