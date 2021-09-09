package routers

import (
	"api_server/api"
	v1 "api_server/api/v1"
	"api_server/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	r.POST("/auth", api.GetAuth)
	r.POST("/login", v1.Login)

	apiv1 := r.Group("/v1")
	{
		apiv1.GET("ocr/result", v1.GeneralBasic)
	}

	auth := r.Group("/v1")
	auth.Use(middleware.JWT())
	{
		auth.GET("/info", v1.GetUserInfo)
	}
	return r
}
