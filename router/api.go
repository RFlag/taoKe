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
}
