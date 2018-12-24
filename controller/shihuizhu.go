package controller

import (
	"github.com/gin-gonic/gin"
	"math"
	"project/dingdangke-dataoke/model"
	"project/ftgo"
)

// 商品类别
func GoodsSort(c *gin.Context) {
	data, err := model.GoodsSort()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

// 商品列表
func GoodsList(c *gin.Context) {
	data, err := model.GoodsList()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}


// 今日疯抢榜
func CrazyToday(c *gin.Context) {
	data, err := model.CrazyToday()
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultData(data))
}

func VideoZone(c *gin.Context) {
	var param struct {
		CateId int `binding:"required" json:"cateId" form:"cateId"`
		Page   int `json:"page" form:"page"`
		Psize  int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Psize <= 0 {
		param.Psize = 10
	}
	data, total, err := model.VideoZone(param.CateId, param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"data":   data,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 排行榜
func Leaderboard(c *gin.Context) {
	var param struct {
		CateId int `binding:"required" json:"cateId" form:"cateId"`
		Page   int `json:"page" form:"page"`
		Psize  int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Psize <= 0 {
		param.Psize = 10
	}
	data, total, err := model.Leaderboard(param.CateId, param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"data":   data,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 头条
func Headline(c *gin.Context) {

}

// 咚咚抢
func Grab(c *gin.Context) {

}

// 半价
func HalfPrice(c *gin.Context) {
	var param struct {
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Psize <= 0 {
		param.Psize = 10
	}
	data, total, err := model.HalfPrice(param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"data":   data,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 9.9
func NinePNine(c *gin.Context) {
	var param struct {
		CateId int `binding:"required" json:"cateId" form:"cateId"`
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Psize <= 0 {
		param.Psize = 10
	}
	data, total, err := model.NinePNine(param.CateId, param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"data":   data,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}

// 推荐
func Recommend(c *gin.Context) {
	var param struct {
		Page  int `json:"page" form:"page"`
		Psize int `json:"psize" form:"psize"`
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Psize <= 0 {
		param.Psize = 10
	}
	data, total, err := model.Recommend( param.Page, param.Psize)
	if err != nil {
		c.AbortWithError(400, err).SetMeta(ResultDrawCouponError).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(200, ftgo.ResultMap(gin.H{
		"data":   data,
		"total":  total,
		"ptotal": math.Ceil(float64(total) / float64(param.Psize)),
	}))
}
