package model

import (
	"dousheng/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Dsn()), &gorm.Config{
		PrepareStmt:                              true,                                // 缓存预编译命令
		SkipDefaultTransaction:                   true,                                // 禁用默认事务操作
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableForeignKeyConstraintWhenMigrating: true,                                // 禁止外键
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
	fmt.Println("InitDB success!")
}
