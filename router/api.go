package router

import (
	"project/dingdangke-dataoke/controller"

	"github.com/gin-gonic/gin"
)

func api(g *gin.RouterGroup) {
	g.POST("/register", controller.Register)

	g.POST("/coupon/draw/list", controller.DrawCouponList)
	g.POST("/popularity/top100/list", controller.Top100PopularityList)
	g.POST("/realtime/amount/list", controller.RealtimeAmountList)
	g.POST("/website/special", controller.WebsiteSpecial)
	g.POST("/qq/special", controller.QqSpecial)
	g.POST("/coupon/draw", controller.DrawCoupon)
	g.POST("/popularity/top100", controller.Top100Popularity)
	g.POST("/realtime/amount", controller.RealtimeAmount)
	g.POST("/goods", controller.Goods)

	g.POST("/shihuizhu/goods/sort", controller.GoodsSort)
	g.POST("/shihuizhu/goods/list", controller.GoodsList)
	g.POST("/shihuizhu/crazy/today", controller.CrazyToday)


	g.POST("/shihuizhu/leaderboard",controller.Leaderboard) // 排行榜
	g.POST("/shihuizhu/headline",controller.Headline) //头条
	g.POST("/shihuizhu/Grab",controller.Grab) //咚咚抢
	g.POST("/shihuizhu/video/zone",controller.VideoZone) // 视频专区
	g.POST("/shihuizhu/half/price",controller.HalfPrice) // 半价
	g.POST("/shihuizhu/nine/point/nine",controller.NinePNine) // 9.9
	g.POST("/shihuizhu/nrecommend",controller.Recommend) // 推荐
}
