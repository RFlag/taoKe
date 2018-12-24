package controller

import (
	"project/dingdangke-dataoke/model"
	"project/ftgo"
	"project/ftgo/ftapi"
	"project/ftgo/ftsql"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var param struct {
		AppId    string `binding:"required" form:"appId" json:"appId"`       // 应用编号
		Platform string `binding:"required" form:"platform" json:"platform"` // 应用平台
		Version  string `binding:"required" form:"version" json:"version"`   // 应用版本
		Captcha  string `binding:"required" form:"captcha" json:"captcha"`   // 验证码
		Nickname string `binding:"required" form:"nickname" json:"nickname"` // 昵称
		Password string `binding:"required" form:"password" json:"password"` // 密码
		Mobile   string `binding:"required" form:"mobile" json:"mobile"`     // 手机号/登录帐号
	}
	err := c.Bind(&param)
	if err != nil {
		return
	}

	err = ftapi.Post("/private/captcha/mobile/check", gin.H{
		"mobile":  param.Mobile,
		"captcha": param.Captcha,
	}, new(ftapi.ResultCodeError))
	if err != nil {
		c.AbortWithError(200, err).SetType(gin.ErrorTypePublic).SetMeta(ResultCaptchaCheckError)
		return
	}
	tx := ftsql.DB.MustBegin()
	userId, err := model.Register(tx, param.Mobile)
	if err != nil {
		c.AbortWithError(200, err).SetType(gin.ErrorTypePublic).SetMeta(ResultAddUserError)
		tx.Rollback()
		return
	}

	var result struct {
		ftapi.ResultCodeError
		Data string
	}
	err = ftapi.Post("/private/register", gin.H{
		"app":      "dataoke",      // 应用编号
		"platform": param.Platform, // 应用平台
		"version":  param.Version,  // 应用版本
		"username": param.Mobile,   // 登录名/昵称
		"password": param.Password, // 密码
		"userId":   userId,         // 用户编号
	}, &result)
	if err != nil {
		c.AbortWithError(200, err).SetType(gin.ErrorTypePublic).SetMeta(ResultCaptchaCheckError)
		tx.Rollback()
		return
	}
	tx.Commit()

	c.JSON(200, ftgo.ResultData(result.Data))
}

// 用户信息
func UserInfo(c *gin.Context)  {

}