package global

import (
	"fmt"
	"time"

	//"github.com/jinzhu/gorm"
    "gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() (error) {
	var err error

	//DBLink, err = gorm.Open(DatabaseSetting.DBType,
	/*
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
	))

	 */
	dsn :=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		DatabaseSetting.UserName,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.DBName,
		DatabaseSetting.Charset,
		DatabaseSetting.ParseTime,)

	DBLink, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	/*
	if ServerSetting.RunMode == "debug" {
		DBLink.LogMode(true)
	}
	DBLink.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DBLink.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	DBLink.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	//otgorm.AddGormCallbacks(db)

	 */
	sqlDB, err := DBLink.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	//DBLink.

	return nil
}
