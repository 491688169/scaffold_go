/*
 * @Author: Kim
 * @Date: 2021-05-13 15:43:37
 * @LastEditTime: 2021-05-13 17:05:36
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/service/service.go
 */
package service

import (
	"context"
	"demo/global"
	"demo/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return &svc
}
