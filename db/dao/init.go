package dao

import (
	"e-learning-platform/config"
	"e-learning-platform/log/logger"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DSN string
)

func InitMySQL() {
	//连接数据库
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Name,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.GormLogger,
	})
	if err != nil {
		panic(errors.New("MySQL初始化错误"))
	}
	//初始化建表
	NewUserDB(db)
}
