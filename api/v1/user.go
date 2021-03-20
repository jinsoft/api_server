package v1

import (
	"github.com/gin-gonic/gin"
	"user_server/request"
)

func Login(c *gin.Context) {
	var param request.Login
	_ = c.ShouldBindJSON(&param)
}
