/*
 * @Author: Kim
 * @Date: 2021-03-08 14:00:59
 * @LastEditTime: 2021-05-17 10:36:12
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/global/setting.go
 */
package global

import (
	"demo/pkg/logger"
	"demo/pkg/setting"

	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS

	DBEngine *gorm.DB
	Logger   *logger.Logger
)
