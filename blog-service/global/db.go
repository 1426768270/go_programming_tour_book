package global

import (
	"blog-service/pkg/setting"
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
var (
	DBEngine *gorm.DB
)

func SetupDBEngine() error {
	var err error
	DBEngine, err = NewDBEngine(DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func NewDBEngine(dbs *setting.DatabaseSettingS)(*gorm.DB, error){
	db, err := gorm.Open(dbs.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbs.UserName,
		dbs.Password,
		dbs.Host,
		dbs.DBName,
		dbs.Charset,
		dbs.ParseTime,
	))
	if err!=nil {
		return nil, err
	}
	if ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 注册3个回调方法 针对我们的公共字段 created_on、modified_on、deleted_on、is_del 进行处理
	db.Callback().Create().Replace("gorm:update_time_stamp",  updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",  updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete",  deleteCallback)
	db.DB().SetMaxIdleConns(dbs.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbs.MaxOpenConns)


	otgorm.AddGormCallbacks(db)
	return db, nil
}
func updateTimeStampForCreateCallback(scope *gorm.Scope){
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok:= scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok:= scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 获取当前设置了标识 gorm:update_column 的字段属性。
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
//  删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 获取当前设置了标识 gorm:delete_option 的字段属性。
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		// 判断是否存在 DeletedOn 和 IsDel 字段，
		// 若存在则调整为执行 UPDATE 操作进行软删除（修改 DeletedOn 和 IsDel 的值），
		// 否则执行 DELETE 进行硬删除。
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
				// 最后在完成一些所需参数设置后调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装。
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