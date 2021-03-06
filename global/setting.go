package global

import (
	"demo/pkg/setting"

	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	DBEngine *gorm.DB
)
