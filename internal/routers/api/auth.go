/*
 * @Author: Kim
 * @Date: 2021-05-12 15:01:22
 * @LastEditTime: 2021-05-13 15:29:05
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/routers/api/auth.go
 */
package api

import (
	"demo/global"
	"demo/internal/service"
	"demo/pkg/app"
	"demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

}
