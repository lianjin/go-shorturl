package storage

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gsurl/log"
)

var db *gorm.DB

type DBLogger struct {
}

func (l *DBLogger) Printf(format string, v ...interface{}) {
	log.Logger.Infof(format, v...)
}

func Init() {
	// 数据库连接参数
	dsn := "dev:123456@tcp(127.0.0.1:3306)/shorturl?charset=utf8mb4&parseTime=True&loc=Local"

	// 日志配置（示例：Info 级别，慢 SQL 超过 1 秒打印，彩色输出）
	gormLogger := logger.New(
		&DBLogger{},
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound错误
			Colorful:                  true,        // 彩色打印
		},
	)

	// 初始化 GORM
	tmpDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Logger.Errorf("failed to connect to database: %v", err)
		panic(err)
	}

	// 设置连接池参数
	sqlDB, err := tmpDB.DB()
	if err != nil {
		log.Logger.Errorf("failed to get db from gorm: %v", err)
		panic(err)
	}
	sqlDB.SetMaxOpenConns(20)                  // 最大连接数
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接最长复用时间
	db = tmpDB
	log.Logger.Info("Database connection initialized successfully")
}
