package load

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

const (
	DefaultMaxLine = 1000
	DefaultMaxSize = 1024
)

/*
[log]
; 日志级别
log_level = 7
; 日志存放目录
log_path = logs/app.log
; 日志文件最大行数
maxlines = 3000
; 日志文件大小限制
maxsize = 10240000
*/

// LoadLoggerConfig 加载默认日志配置
func LoadLoggerConfig(LogPath string) {
	logs.Info("----> load logger config from: %s", LogPath)
	BConfig, err := config.NewConfig("ini", LogPath)
	if err != nil {
		fmt.Println("config init error:", err)
		return
	}
	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = DefaultMaxLine
	}
	maxsize, serr := BConfig.Int64("log::maxsize")
	if serr != nil {
		maxsize = DefaultMaxSize
	}
	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::log_path")
	level, _ := BConfig.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxlines
	logConf["maxsize"] = maxsize

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}

	_ = logs.SetLogger(logs.AdapterFile, string(confStr))
	logs.SetLogFuncCall(true)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Info("----> load logger config success.")
}
