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

	r.POST("/api/r1/query/user/info", release1.QueryUserInfo)
	r.POST("/api/r1/query/user/credit_balance", release1.QueryBalance)
	r.POST("/api/r1/query/block/rawdata", release1.QueryRawBlock)


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
