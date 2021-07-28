package r1

import (
	"github.com/jack139/artchain/cmd/http/helper"
	"github.com/jack139/artchain/cmd/ipfs"

	"github.com/valyala/fasthttp"
	"log"
)

/* 从ipfs下载数据 */
func IpfsDownload(ctx *fasthttp.RequestCtx) {
	log.Println("ipfs_download")

	// POST 的数据
	content := ctx.PostBody()

	// 验签
	reqData, err := helper.CheckSign(content)
	if err != nil {
		helper.RespError(ctx, 9000, err.Error())
		return
	}

	// 检查参数
	hash, ok := (*reqData)["hash"].(string)
	if !ok {
		helper.RespError(ctx, 9001, "need hash")
		return
	}

	if len(hash) == 0 {
		helper.RespError(ctx, 9002, "need hash")
		return
	}

	var image_data []byte

	// 检查redis
	image_data, err = helper.GetImage(hash)
	if err != nil {
		helper.RespError(ctx, 9004, err.Error())
		return
	}

	if image_data == nil {
		// 从ipfs读取
		image_data, err = ipfs.Get(hash)
		if err != nil {
			helper.RespError(ctx, 9005, err.Error())
			return
		}

		// 缓存到 redis
		err = helper.CacheImage(hash, image_data)
		if err != nil {
			helper.RespError(ctx, 9006, err.Error())
			return
		}
		log.Println("Cache image: ", hash)
	}

	// 生成返回数据
	resp := map[string]interface{}{
		"data": string(image_data),
	}

	helper.RespJson(ctx, &resp)
}
