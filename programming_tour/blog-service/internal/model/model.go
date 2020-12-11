package model

import (
	"fmt"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"   // MySQL 驱动库初始化
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

/*
问题：接口返回，很多字段的值都是零值呢？ 可以发现有值的部分都是筛选（Select）出来的字段，因而出现问题的原因是，没有取值的字段也显示了出来（默认取零值）

解决：
1. 使用 omitempty
2. model.Model 是公共 model，如果对它进行修改，岂不是所有的代码都会受影响。特别是，对于某些字段来说，零值就是默认值，如果直接修改，则影响还是很大的。
   因而在追求隔离性的情况下，我们可以重新定义一个结构体，用于处理该Service方法的结果集的返回
 */

// Json 关键字 omitempty：
// 当字段的tag规则包含omitempty属性时，如果该字段的值为该字段类型的零值，那么在进行转换时会忽略该字段。
// 简单来说，就是在设置 omitempty 后，如果这个字段的值为零值，则不显示


// 公共 model
type Model struct {

	ID         uint32 `gorm:"primary_key" json:"id"`
	//ID         uint32 `gorm:"primary_key" json:"id,omitempty"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 采用设置model callback的方式实现对公共字段的处理
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	otgorm.AddGormCallbacks(db)
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 通过调用scope.FieldByName方法，获取当前是否包含所需的字段。
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// 通过判断Field.IsBlank的值，可以得知该字段的值是否为空
			// 若为空，则调用Field.Set方法给该字段设置值。入参类型为interface{}，即内部是通过反射进行一系列操作赋值的
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
	// 通过调用 scope.Get（″gorm：update_column″） 来获取当前设置的标识gorm：update_column的字段属性
	// 若不存在，即没有自定义设置 update_column，则在更新回调内设置默认字段ModifiedOn的值为当前的时间戳
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 通过调用scope.Get（″gorm：delete_option″）来获取当前设置的标识gorm：delete_option的字段属性
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// 判断是否存在DeletedOn和IsDel字段。
		// 若存在，则执行UPDATE操作进行软删除（修改DeletedOn和IsDel的值），否则执行DELETE操作进行硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
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
		/*
		调用scope.QuotedTableName方法获取当前引用的表名，并调用一系列方法对SQL语句的组成部分进行处理和转移。
		在完成一些所需参数设置后，调用scope.CombinedConditionSql方法完成SQL语句的组装
		 */
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
