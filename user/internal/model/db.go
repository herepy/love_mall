/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/9/6 18:07
 */

package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"love_mall/user/internal/config"
	"os"
	"time"
)

var db *gorm.DB

func InitDb(cfg config.Config) {
	if db != nil {
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User,
		cfg.Mysql.Pwd,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Db)

	var err error
	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second * 2, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,     // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(3)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
}

func GetDB() *gorm.DB {
	return db
}
