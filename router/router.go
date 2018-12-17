package router

import "github.com/gin-gonic/gin"

func Router(g *gin.Engine) {
	api(g.Group("/api"))
	manager(g.Group("/manager"))
}
