package middleware

import (
	"strings"

	"project/login/lib/jwt"

	"project/ftgo"
	"project/ftgo/ftapi"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	PayloadKey = "payload"
	UserKey    = "user"
)

var (
	ResultCheckTokenError = ftgo.ResultError("验证token失败")
)

func CheckToken(appId, appPlatform, appVersion, action string, getUser func(payload *jwt.Payload) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			c.AbortWithError(400, err).SetType(gin.ErrorTypeBind)
			return
		}

		result := new(struct {
			ftapi.ResultCodeError
			Data *jwt.Payload
		})
		err = ftapi.Post("/private/token/check", map[string]string{
			"app":      appId,
			"platform": appPlatform,
			"version":  appVersion,
			"action":   action,
			"token":    token,
		}, result)
		if err != nil {
			c.AbortWithError(401, errors.WithMessage(err, "检查 token 失败")).SetMeta(ResultCheckTokenError).SetType(gin.ErrorTypePublic)
			return
		}
		c.Set(PayloadKey, result.Data)

		if getUser != nil {
			userInfo, err := getUser(result.Data)
			if err != nil {
				c.AbortWithError(200, errors.WithMessage(err, "获取用户信息失败")).SetMeta(ResultCheckTokenError).SetType(gin.ErrorTypePublic)
				return
			}
			c.Set(UserKey, userInfo)
		}
	}
}

// 获取 token 参数
func getToken(c *gin.Context) (string, error) {
	if Authorization := c.GetHeader("Authorization"); strings.HasPrefix(Authorization, "Bearer ") {
		return strings.TrimPrefix(Authorization, "Bearer "), nil
	} else {
		t := c.Query("access_token")
		if t == "" {
			return "", errors.New("没有 token 参数")
		}
		return t, nil
	}
}
