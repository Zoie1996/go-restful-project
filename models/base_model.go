package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"go-restful-project/pkg/setting"
	"go-restful-project/pkg/util"
	"time"
)

var db *gorm.DB

type BaseModel struct {
	ID         int           `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Status     int8          `gorm:"column:status;default:1" json:"status"`
	IsDel      bool          `gorm:"column:is_del;default:false" json:"is_del"`
	CreateTime util.JSONTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime util.JSONTime `gorm:"column:update_time" json:"update_time"`
	DelTime    time.Time     `gorm:"column:del_time;default:null" json:"del_time"`
}

// 安装程序初始化数据库实例
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Password))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB关闭数据库连接
func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback将在创建时设置“CreateTime”、“UpdateTime”
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback将在更新时设置“UpdateTime”
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", time.Now())
	}
}

// deleteCallback将在删除的地方设置“IsDel”、“DelTime”
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		delTimeField, hasDelTimeField := scope.FieldByName("DelTime")

		if !scope.Search.Unscoped && hasIsDelField && hasDelTimeField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(true),
				scope.Quote(delTimeField.DBName),
				scope.AddToVars(time.Now()),
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

// addExtraSpaceIfExist添加了一个分隔符
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
