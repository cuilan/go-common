package test

import (
	"github.com/cuilan/go-common/orm"
	"testing"
)

func TestInitMysql(t *testing.T) {
	orm.InitMysql("weattech:E1W(kF%C_X=W7FO~hneGDh5wH+vIJ2Ia@tcp(10.110.1.143:5807)/download_db?parseTime=True&loc=Local&charset=utf8mb4", 1, 1)
}
