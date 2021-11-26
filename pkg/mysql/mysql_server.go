package mysql

import (
	"fmt"
	"frame/consts"
	"frame/global"
	gormPlugin "frame/pkg/jaeger"
	"frame/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

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
		db.LogMode(false)
	}

	// 默认使用单表
	db.SingularTable(true)

	// 空闲连接最大连接数
	db.DB().SetMaxIdleConns(setting.MaxIdleConns)
	// 最大打开连接数
	db.DB().SetMaxOpenConns(setting.MaxOpenConns)
	// 设置最大连接的时间
	db.DB().SetConnMaxLifetime(time.Second * 900) // 15 分钟
	// 设置最大的时间闲置的连接
	db.DB().SetConnMaxIdleTime(time.Second * 10) // 10 秒钟

	// 增加 openTracing 回调
	gormPlugin.AddGormCallbacks(db)

	return db, nil
}

// ----------------------------------------------注册回调函数--------------------------------------
//db.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
//db.Callback().Create().Replace("gorm:delete", deleteCallback)
// -----------------------------------------------注册回调函数--------------------------------------
// 再编写model的时候 并没有对 CreatedOn ModifiedOn DeletedOn IsDel 进行处理
// 再每张表中都插入这些字段显然不是很好的选择
// 我们采用model callback 的方式进行处理,GORM本身也支持回调
/*
	1) 注册一个新的回调
	2) 删除现有的回调
	3) 替换现有的回调
	4) 注册回调的先后顺序
*/

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	/*
		1) 通过调用 scope.FieldByName, 获取当前是否包含所需的字段
		2) 通过判断 Field.IsBlank的值, 可以得知字段的值是否为空
		3) 为空传递默认值, 参数格式为 interface, 内部通过放射来获取值
	*/
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	/*
		1) 通过  scope.Get("gorm:update_column") 获取更新的时间
		2) 如果不存再则设置默认值
	*/
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDeletedField, hasIsDeletedField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDeletedField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDeletedField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
