/*
 * @Author: Kim
 * @Date: 2021-03-09 17:39:13
 * @LastEditTime: 2021-03-09 17:41:15
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /scaffold_go/internal/service/tag.go
 */
package service

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
