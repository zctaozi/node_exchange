package routers

import (
	"bytes"
	"chainup.com/node-exchange/middlewares"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
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
	router.POST("/", proxy)
}

func proxy(c *gin.Context) {
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	request, err := http.NewRequest(http.MethodPost, "https://smart.zeniq.network:9545", bytes.NewReader(reqBody))
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	response, err := client.Do(request)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}
	c.String(http.StatusOK, string(responseBody))
}
