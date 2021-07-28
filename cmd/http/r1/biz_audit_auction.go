package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

/* 审核拍卖 */
func BizAuditAuction(ctx *fasthttp.RequestCtx) {
	log.Println("biz_audit_auction")

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
	auctionIdStr, ok := (*reqData)["id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need id")
		return
	}
	status, ok := (*reqData)["status"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need status")
		return
	}
	openDate, _ := (*reqData)["open_date"].(string)
	closeDate, _ := (*reqData)["close_date"].(string)

	auctionId, err := strconv.ParseUint(auctionIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	auctionMap, err := queryAuctionInfoById(auctionId)
	if err != nil {
		helper.RespError(ctx, 9002, err.Error())
		return
	}

	// 检查是否修改拍卖时间
	if len(openDate) == 0 {
		openDate = (*auctionMap)["openDate"].(string)
	}
	if len(closeDate) == 0 {
		closeDate = (*auctionMap)["closeDate"].(string)
	}

	// 修改链上数据
	respData, err := auctionModify(auctionMap, callerAddr, auctionId,
		"\x00", "\x00", openDate, closeDate, status, "audit")
	if err != nil {
		helper.RespError(ctx, 9010, err.Error())
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"height": (*respData)["height"].(string), // 区块高度
	}

	helper.RespJson(ctx, &resp)
}
