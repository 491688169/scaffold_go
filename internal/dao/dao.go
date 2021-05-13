/*
 * @Author: Kim
 * @Date: 2021-05-13 15:45:07
 * @LastEditTime: 2021-05-13 15:47:55
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/dao/dao.go
 */
package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{
		engine: engine,
	}
}
