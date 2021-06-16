package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"

	"log"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

/* 删除物品图片 - 并为实际从ipfs删除，只是从链上数据中删除 */

func IpfsRemoveImage(ctx *fasthttp.RequestCtx) {
	log.Println("ipfs_remove_image")

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
	itemIdStr, ok := (*reqData)["item_id"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need item_id")
		return
	}

	hash, ok := (*reqData)["hash"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need hash")
		return
	}

	itemId, err := strconv.ParseUint(itemIdStr, 10, 64)
	if err != nil {
		helper.RespError(ctx, 9007, err.Error())
		return
	}

	// 获取当前链上数据
	itemMap, err := queryItemInfoById(itemId)
	if err!=nil {
		helper.RespError(ctx, 9002, err.Error())
		return		
	}

	// 准备数据
	loadData := (*itemMap)["itemImage"].([]string)	
	//log.Printf("1: %v", loadData)
	for i, h := range loadData{
		if h==hash { // 删除 图片的  hash
			loadData = append(loadData[:i], loadData[i+1:]...)
			break
		}
	}
	//log.Printf("2: %v", loadData)
	loadBytes, err := json.Marshal(loadData)
	if err != nil {
		helper.RespError(ctx, 9008, err.Error())
		return
	}

	// 修改链上数据
	respData, err := itemModify(itemMap, callerAddr, 
		itemId, "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", 
		string(loadBytes), "\x00", "\x00", "\x00", "WAIT", "remove image")
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
