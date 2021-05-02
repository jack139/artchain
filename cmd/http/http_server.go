package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"

	"github.com/jack139/artchain/cmd/http/helper"
	release1 "github.com/jack139/artchain/cmd/http/r1"
)


/* 入口 */
func RunServer(port string /*, userPath string*/) {

	/* router */
	r := router.New()
	r.GET("/", index)
	r.POST("/api/test", doNonthing)

	r.POST("/api/r1/biz/user/register", release1.BizRegister)
	r.POST("/api/r1/biz/user/modify", release1.BizUserModify)
	r.POST("/api/r1/biz/item/new", release1.BizItemNew)
	r.POST("/api/r1/biz/item/modify", release1.BizItemModify)
	r.POST("/api/r1/biz/review/new", release1.BizReviewNew)
	r.POST("/api/r1/biz/review/modify", release1.BizReviewModify)
	r.POST("/api/r1/biz/auction/new", release1.BizAuctionNew)
	r.POST("/api/r1/biz/auction/modify", release1.BizAuctionModify)
	r.POST("/api/r1/biz/trans/new", release1.BizTransNew)

	r.POST("/api/r1/query/user/info", release1.QueryUserInfo)
	r.POST("/api/r1/query/user/credit_balance", release1.QueryBalance)
	r.POST("/api/r1/query/block/rawdata", release1.QueryRawBlock)
	r.POST("/api/r1/query/item/info", release1.QueryItemInfo)
	r.POST("/api/r1/query/item/list", release1.QueryItemList)
	r.POST("/api/r1/query/review/list", release1.QueryReviewList)
	r.POST("/api/r1/query/auction/info", release1.QueryAuctionInfo)
	r.POST("/api/r1/query/auction/list", release1.QueryAuctionList)
	r.POST("/api/r1/query/trans/info", release1.QueryTransInfo)
	r.POST("/api/r1/query/trans/list", release1.QueryTransList)


	log.Printf("start HTTP server at 0.0.0.0:%s\n", port)

	/* 启动server */
	s := &fasthttp.Server{
		Handler: helper.Combined(r.Handler),
		Name:    "FastHttpLogger",
	}
	log.Fatal(s.ListenAndServe(":" + port))
}

/* 根返回 */
func index(ctx *fasthttp.RequestCtx) {
	log.Printf("%v", ctx.RemoteAddr())
	ctx.WriteString("Hello world.")
}
