/*
 * @Author: Kim
 * @Date: 2021-03-08 16:51:16
 * @LastEditTime: 2021-03-08 17:04:34
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /scaffold_go/pkg/app/application.go
 */
package app

import (
	"demo/global"
	"demo/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) int {
	page := convert.StrTo(ctx.Query("page")).MustToInt()

	if page < 0 {
		return 1
	}

	return page
}

func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustToInt()

	if pageSize < 0 {
		pageSize = global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		pageSize = global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0

	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
