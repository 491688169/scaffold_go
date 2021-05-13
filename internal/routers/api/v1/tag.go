/*
 * @Author: Kim
 * @Date: 2021-03-08 17:50:00
 * @LastEditTime: 2021-05-12 10:25:25
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/routers/api/v1/tag.go
 */

package v1

import (
	"demo/pkg/app"
	"demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
