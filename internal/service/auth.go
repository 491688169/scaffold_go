/*
 * @Author: Kim
 * @Date: 2021-05-12 15:10:01
 * @LastEditTime: 2021-05-13 17:39:07
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/service/auth.go
 */
package service

import (
	"errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)

	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
