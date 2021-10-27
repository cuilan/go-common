package orm

import (
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitMysql 初始化MySQL连接
// Dsn 数据库连接地址
// MysqlMaxActive
// MysqlMaxIdle
func InitMysql(Dsn string, MysqlMaxActive, MysqlMaxIdle int) (Db *gorm.DB) {
	var err error
	logs.Info("开始初始化数据库")
	// 定义 sql 日志类型
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	// "%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&charset=%s"
	logs.Info(Dsn)

	Db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database.")
	}

	sqlDB, serr := Db.DB()
	if serr != nil {
		panic("failed open sql db pool.")
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(MysqlMaxIdle)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(MysqlMaxActive)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()
	if err != nil {
		logs.Error("connect mysql err:" + err.Error())
	}
	logs.Info("初始化数据库完成")
	return Db
}
