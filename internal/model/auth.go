/*
 * @Author: Kim
 * @Date: 2021-03-09 11:32:30
 * @LastEditTime: 2021-05-12 11:50:15
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/model/auth.go
 */
package model

import "gorm.io/gorm"

type Auth struct {
	*Model
	Appkey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a *Auth) TableName() string {
	return "demo_auth"
}

func (a *Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.Appkey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
