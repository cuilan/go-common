package nsq

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/nsqio/go-nsq"
	"os"
)

// Producer nsq发布消息
func Producer(msgBody []byte, topic string) {
	// 新建生产者
	producer, err := nsq.NewProducer(NsqHost, nsq.NewConfig())
	if err != nil {
		logs.Error("新建生产者", err)
	}
	err = producer.Ping()
	if err != nil {
		logs.Error("p.Ping()", err)
	}
	// 发布消息
	if err := producer.Publish(topic, msgBody); err != nil {
		logs.Error("发布消息", err)
		// TODO 业务是否需要判断退出
		os.Exit(0)
	} else {
		logs.Info("发布成功")
	}
	producer.Stop()
}

// AddNsq 将结果信息添加到订阅信息中 topic
func AddNsq(file interface{}) error {
	dataByte, err := json.Marshal(file)
	if err != nil {
		return err
	}
	Producer(dataByte, NsqTopicName)
	return nil
}
