package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const defaultDatabaseName = "demo"

var DB *gorm.DB

func SetupFromEnv() {
	DB = setupMySQL()
}

func Close() {
	// 测试发现在程序重启时 sqlite 似乎并不积极地合并 shm 和 wal 文件到 db 文件
	// 在程序退出主动关闭一下数据库连接可能更好
	if DB != nil {
		if conn, err := DB.DB(); err == nil {
			_ = conn.Close()
		}
	}
}

func init() {
	showSQL := os.Getenv("DATABASE_SHOW_SQL") == "true"
	conf := gormLogger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  gormLogger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
	if showSQL {
		conf.LogLevel = gormLogger.Info
	}

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(logrus.StandardLogger().Formatter)
	// 所有 sql 语句输出到 info 级别 logrus 中
	gormLogger.Default = gormLogger.New(log.New(logger.Writer(), "\r\n", log.LstdFlags), conf)
}
