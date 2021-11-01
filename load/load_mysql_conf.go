package load

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/go-common/orm"
	"gorm.io/gorm"
)

/*
[mysql]
host=localhost
user=root
password=123456
port=3306
db=weattech_dl
charset=utf8mb4
maxactive=10
maxidle=10

[mysql-xxx] 多数据源名称
*/

// LoadMysqlConfig 初始化 MySQL 配置信息
func LoadMysqlConfig(ConfigPath string) (Db *gorm.DB) {
	dsn, mysqlMaxactive, mysqlMaxidle := LoadMysqlWithDataSource(ConfigPath, "mysql")
	return orm.InitMysql(dsn, mysqlMaxactive, mysqlMaxidle)

}

// LoadMysqlWithDataSource 读取 mysql 配置
// ConfigPath 配置文件路径
// DataSourceName 数据源配置名称
func LoadMysqlWithDataSource(ConfigPath, DataSourceName string) (dsn string, mysqlMaxactive, mysqlMaxidle int) {
	logs.Info("----> load mysql config from: %s", ConfigPath)
	Cfg, err := config.NewConfig("ini", ConfigPath)
	if err != nil {
		logs.Info("config init error:", err)
		panic(err)
	}
	mysqlHost := Cfg.String(DataSourceName + "::host")
	mysqlUser := Cfg.String(DataSourceName + "::user")
	mysqlPassword := Cfg.String(DataSourceName + "::password")
	mysqlPort, _ := Cfg.Int(DataSourceName + "::port")
	mysqlDb := Cfg.String(DataSourceName + "::db")
	mysqlCharset := Cfg.String(DataSourceName + "::charset")
	mysqlMaxactive, _ = Cfg.Int(DataSourceName + "::maxactive")
	mysqlMaxidle, _ = Cfg.Int(DataSourceName + "::maxidle")
	dsn = FormatDsn(mysqlHost, mysqlUser, mysqlPassword, mysqlDb, mysqlCharset, mysqlPort)
	logs.Info("----> load mysql config success.")
	return dsn, mysqlMaxactive, mysqlMaxidle
}

// FormatDsn 格式化数据库连接 DSN
func FormatDsn(MysqlHost, MysqlUser, MysqlPassword, MysqlDb, MysqlCharset string, MysqlPort int) (dsn string) {
	formatStr := "%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local&charset=%s"
	return fmt.Sprintf(formatStr, MysqlUser, MysqlPassword, MysqlHost, MysqlPort, MysqlDb, MysqlCharset)
}
