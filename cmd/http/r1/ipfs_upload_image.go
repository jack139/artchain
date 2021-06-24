package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	"github.com/jack139/artchain/cmd/ipfs"

	"log"
	"strconv"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

/* 上传物品图片 */

func IpfsUploadImage(ctx *fasthttp.RequestCtx) {
	log.Println("ipfs_upload_image")

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

	image, ok := (*reqData)["image"].(string)
	if !ok {
		helper.RespError(ctx, 9002, "need image")
		return
	}

	if len(image)>5242880 {
		helper.RespError(ctx, 90011, "image too large: over 5M")
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


	// 图片 存 ipfs
	var cid string
	if len(image)>0 {
		cid, err = ipfs.Add([]byte(image))
		if err!=nil {
			helper.RespError(ctx, 9012, err.Error())
			return
		}
	} else {
		cid = ""
	}

	// 准备数据
	loadData := (*itemMap)["itemImage"].([]string)	
	loadData = append(loadData, cid)
	loadBytes, err := json.Marshal(loadData)
	if err != nil {
		helper.RespError(ctx, 9008, err.Error())
		return
	}

	// 修改链上数据
	respData, err := itemModify(callerAddr, 
		itemId, "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", "\x00", 
		string(loadBytes), "\x00", "\x00", "\x00", "WAIT", "upload image")
	if err != nil {
		helper.RespError(ctx, 9010, err.Error())
		return
	}

	// 返回区块id
	resp := map[string]interface{}{
		"height" : (*respData)["height"].(string),  // 区块高度
		"hash"   : cid, // ipfs hash
	}

	helper.RespJson(ctx, &resp)
}
