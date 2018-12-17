package controller

import (
	"math"

	"project/dingdangke-dataoke/model"
	"project/ftgo"

	"github.com/gin-gonic/gin"
)

// 网站专用API接口
func WebsiteSpecial(c *gin.Context) {
	data, err := model.WebsiteSpecial()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// qq群发专用API接口
func QqSpecial(c *gin.Context) {
	data, err := model.QqSpecial()
	if err != nil {
		c.AbortWithError(200, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// 全站领券商品API接口
func DrawCoupon(c *gin.Context) {
	data, err := model.DrawCoupon()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// 领券商品列表
func DrawCouponList(c *gin.Context) {
	var param struct {
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	result, total, err := model.DrawCouponList(param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"result": result,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// TOP100人气榜API接口
func Top100Popularity(c *gin.Context) {
	data, err := model.Top100Popularity()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// TOP100人气榜列表
func Top100PopularityList(c *gin.Context) {
	var param struct {
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	result, total, err := model.Top100PopularityList(param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"result": result,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 实时跑量榜API接口
func RealtimeAmount(c *gin.Context) {
	data, err := model.RealtimeAmount()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// 实时跑量榜列表
func RealtimeAmountList(c *gin.Context) {
	var param struct {
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	result, total, err := model.RealtimeAmountList(param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"result": result,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 单品详情API接口
func Goods(c *gin.Context) {
	var param struct {
		Id int `json:"id" form:"id"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	data, err := model.Goods(param.Id)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}
