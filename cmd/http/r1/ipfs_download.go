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

	// 处理image 字段，从ipfs读取
	var image_data []byte
	if len(hash)>0 {
		image_data, err = ipfs.Get(hash)
		if err!=nil {
			helper.RespError(ctx, 9005, err.Error())
			return
		}
	}

	//log.Printf("%v\n", string(respBytes))

	// 生成返回数据
	resp := map[string]interface{}{
		"data": string(image_data),
	}

	helper.RespJson(ctx, &resp)
}

