package v1

import (
	"api_server/pkg/app"
	"api_server/request"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var param request.Login
	_ = c.ShouldBindJSON(&param)
}

func GetUserInfo(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"name": "hello world",
	})
}
