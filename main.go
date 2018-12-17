package main

import (
	"project/dingdangke-dataoke/router"
	"project/ftgo"

	"github.com/gin-gonic/gin"
)

var version string
var port = "80"

func main() {
	ftgo.Run(":"+port, func(g *gin.Engine) {
		g.Any("/version", func(c *gin.Context) {
			c.String(200, version)
		})
		router.Router(g)
	})
}
