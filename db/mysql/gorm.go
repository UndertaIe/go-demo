package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewDBEngine() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:mysql@tcp(127.0.0.1:3306)/go_gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                          // 根据版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db.Debug(), nil // Debug可以打印调用方法和sql语句
}

func init() {
	var err error
	Db, err = NewDBEngine()
	if err != nil {
		panic("init mysql error")
	}
}
