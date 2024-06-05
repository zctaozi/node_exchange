package routers

import (
	"chainup.com/node-exchange/api"
	"chainup.com/node-exchange/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.NoRoute(NoResponse)
	engine.Use(middlewares.Cors())
	initRouter(engine)
	return engine
}

func NoResponse(c *gin.Context) {
	// 返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}

func initRouter(router *gin.Engine) {
	group := router.Group("/api")
	group.POST("/getTx", api.GetRecipet)
}
