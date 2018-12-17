package controller

import (
	"project/ftgo"
)

var (
	ResultDrawCouponError     = ftgo.ResultError("全站领券接口失败")
	ResultDrawCouponListError = ftgo.ResultError("获取领券列表失败")
	ResultCaptchaCheckError   = ftgo.ResultError("手机验证码验证失败")
	ResultRegisterError       = ftgo.ResultError("注册失败")
	ResultAddUserError        = ftgo.ResultError("添加用户失败")
)
