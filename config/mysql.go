package config

import (
	"log"
	"time"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlProviderSet = wire.NewSet(NewMysql)

func NewMysql(AppConfig *Config) *gorm.DB {
	if AppConfig.Mysql.Dsn != "" {
		db, err := gorm.Open(mysql.Open(AppConfig.Mysql.Dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect mysql database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get database instance: %v", err)
		}
		//设置数据库连接池最大空闲数量
		sqlDB.SetMaxIdleConns(AppConfig.Mysql.MaxIdleConn)
		//设置数据库连接池中打开的最大连接数
		sqlDB.SetMaxOpenConns(AppConfig.Mysql.MaxOpenConn)
		//设置数据库连接的最大生命周期
		sqlDB.SetConnMaxLifetime(time.Hour)

		return db
	}
	return nil
}
