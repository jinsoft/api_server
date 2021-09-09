package routers

import (
	"api_server/api"
	v1 "api_server/api/v1"
	_ "api_server/docs" // you have to import it.
	"api_server/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	url := ginSwagger.URL("http://127.0.0.1:8300/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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
