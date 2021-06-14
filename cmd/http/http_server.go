package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"time"

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
	r.POST("/api/r1/biz/auction/bid", release1.BizAuctionBid)
	r.POST("/api/r1/biz/auction/bid/withdraw", release1.BizAuctionBidWithdraw)
	r.POST("/api/r1/biz/trans/new", release1.BizTransNew)
	r.POST("/api/r1/biz/audit/user", release1.BizAuditUser)
	r.POST("/api/r1/biz/audit/item", release1.BizAuditItem)
	r.POST("/api/r1/biz/audit/review", release1.BizAuditReview)
	r.POST("/api/r1/biz/audit/auction", release1.BizAuditAuction)

	r.POST("/api/r1/query/user/info", release1.QueryUserInfo)
	r.POST("/api/r1/query/user/list", release1.QueryUserList)
	r.POST("/api/r1/query/user/verify", release1.QueryUserVerify)
	r.POST("/api/r1/query/user/credit_balance", release1.QueryBalance)
	r.POST("/api/r1/query/block/rawdata", release1.QueryRawBlock)
	r.POST("/api/r1/query/item/info", release1.QueryItemInfo)
	r.POST("/api/r1/query/item/list", release1.QueryItemList)
	r.POST("/api/r1/query/review/info", release1.QueryReviewInfo)
	r.POST("/api/r1/query/review/list", release1.QueryReviewList)
	r.POST("/api/r1/query/auction/info", release1.QueryAuctionInfo)
	r.POST("/api/r1/query/auction/list", release1.QueryAuctionList)
	r.POST("/api/r1/query/trans/info", release1.QueryTransInfo)
	r.POST("/api/r1/query/trans/list", release1.QueryTransList)
	r.POST("/api/r1/query/bid/info", release1.QueryBidInfo)
	r.POST("/api/r1/query/bid/highest", release1.QueryBidHighest)
	r.POST("/api/r1/query/bid/list", release1.QueryBidList)
	r.POST("/api/r1/query/auction_house/list", release1.QueryAHList)
	r.POST("/api/r1/query/user/list_by_status", release1.QueryUserListByStatus)
	r.POST("/api/r1/query/item/list_by_status", release1.QueryItemListByStatus)
	r.POST("/api/r1/query/review/list_by_status", release1.QueryReviewListByStatus)
	r.POST("/api/r1/query/auction/list_by_status", release1.QueryAuctionListByStatus)

	r.POST("/api/r1/ipfs/upload/image", release1.IpfsUploadImage)
	r.POST("/api/r1/ipfs/download", release1.IpfsDownload)
	r.POST("/api/r1/ipfs/remove/image", release1.IpfsRemoveImage)



	// 设置定时任务 : 15 秒 一次
	ticker1 := time.NewTicker(5 * time.Second)
	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			// 检查拍卖时间，进行状态转换
			if err := release1.CheckAuction(); err!=nil{
				log.Println(err.Error())
			}
			log.Println("--> Ticker")
		}
	}(ticker1)


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
