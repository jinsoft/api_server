package api

import (
	"api_server/global"
	"api_server/internal/service"
	"api_server/pkg/app"
	"api_server/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.LOG.Error("app.BindAndValid errs")
		if errs != nil {
			global.LOG.Error("")
		}
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("check auth error: %v", err.Error()))
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("generate token error: %v", err.Error()))
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
