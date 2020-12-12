package models

import (
	"gobot/pkg/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// DB 数据库链接单例
var db *gorm.DB

func Init() {
	log.Info("初始化数据库")
	var err error

	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("连接数据库不成功: %s", err)

	}
	config()
	migrate()
}

func config() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("数据库配置失败: %s", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// Migrate the schema
func migrate() {
	err := db.AutoMigrate(
		&Task{},
	)
	if err != nil {
		log.Panic(err)
	}
}
