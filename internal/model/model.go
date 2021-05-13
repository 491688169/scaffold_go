/*
 * @Author: Kim
 * @Date: 2021-03-08 14:00:59
 * @LastEditTime: 2021-05-12 11:18:33
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/internal/model/model.go
 */
package model

import (
	"demo/global"
	"demo/pkg/setting"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	gorm.Model
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	sqldb, err := db.DB()
	sqldb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqldb.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil

}
