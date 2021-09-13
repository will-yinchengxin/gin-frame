package model

import (
	"fmt"
	"frame/consts"
	"frame/global"
	"frame/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	CreateOn   int64  `gorm:"column:create_on" json:"create_on" form:"create_on"`
	CreateBy   string `gorm:"column:create_by" json:"create_by" form:"create_by"`
	ModifyedOn int64  `gorm:"column:modifyed_on" json:"modifyed_on" form:"modifyed_on"`
	ModifyedBy string `gorm:"column:modifyed_by" json:"modifyed_by" form:"modifyed_by"`
	DeletedOn  int64  `gorm:"column:deleted_on" json:"deleted_on" form:"deleted_on"`
	IsDel      int64  `gorm:"column:is_del" json:"is_del" form:"is_del"`
}

func NewDBEngine(setting *setting.DatabaseSetting) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		setting.Username,
		setting.Password,
		setting.Host,
		setting.DBName,
		setting.Charset,
		setting.ParseTime,
	)
	db, err := gorm.Open(setting.DBType, dns)
	if err != nil {
		return nil, err
	}
	// 开发模式开启日志详细模式
	if global.ServerSetting.RunMode == consts.RunMode {
		db.LogMode(true)
	}
	// 默认使用单表
	db.SingularTable(true)
	// 空闲连接最大连接数
	db.DB().SetMaxIdleConns(setting.MaxIdleConns)
	// 最大打开连接数
	db.DB().SetMaxOpenConns(setting.MaxOpenConns)

	return db, nil
}
