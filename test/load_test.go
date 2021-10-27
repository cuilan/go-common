package test

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-common/v0.0.1/entity"
	"github.com/go-common/v0.0.1/load"
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

	var dlRecords = &entity.DlRecords{}
	result := ds.Model(&entity.DlRecords{}).Where("`desc` != ?", "").Scan(dlRecords)
	if result.Error != nil {
		logs.Info(result.Error)
		panic(result.Error)
	}
	fmt.Println(dlRecords)
}

// 测试读取自定义数据源
func TestLoadMysqlWithDataSource(t *testing.T) {
	source, maxactive, maxidle := load.LoadMysqlWithDataSource("../conf/mysql.conf", "mysql-haulers")
	fmt.Println(source)
	fmt.Println(maxactive)
	fmt.Println(maxidle)
}
