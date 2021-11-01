package test

import (
	"fmt"
	"github.com/cuilan/go-common/load"
	"gorm.io/gorm"
	"testing"
)

func TestLoadLoggerConfig(t *testing.T) {
	load.LoadLoggerConfig("../conf/log.conf")
}

func TestLoadMinioConfig(t *testing.T) {
	load.LoadMinioConfig("../conf/minio.conf")
}

func TestLoadMysqlConfig(t *testing.T) {
	var ds *gorm.DB
	ds = load.LoadMysqlConfig("../conf/mysql.conf")
	fmt.Println(ds)
}

// 测试读取自定义数据源
func TestLoadMysqlWithDataSource(t *testing.T) {
	source, maxactive, maxidle := load.LoadMysqlWithDataSource("../conf/mysql.conf", "mysql-haulers")
	fmt.Println(source)
	fmt.Println(maxactive)
	fmt.Println(maxidle)
}
