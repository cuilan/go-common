package load

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/go-common/nsq"
)

/*
[nsq]
host=127.0.0.1
topic_name=
*/

// LoadNsqConfig 加载 NSQ 配置信息
func LoadNsqConfig(ConfigPath string) {
	logs.Info("---->  load nsq config from: %s", ConfigPath)
	Cfg, err := config.NewConfig("ini", ConfigPath)
	if err != nil {
		logs.Info("config init error:", err)
		panic(err)
	}
	nsq.NsqHost = Cfg.String("nsq::host")
	nsq.NsqTopicName = Cfg.String("nsq::topic_name")
	logs.Info("----> load nsq config success.")
}
