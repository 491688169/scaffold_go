/*
 * @Author: Kim
 * @Date: 2021-05-13 17:17:05
 * @LastEditTime: 2021-05-13 17:23:19
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/dao/auth.go
 */
package dao

import "demo/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{Appkey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
