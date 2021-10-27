package load

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var MinioConfig MinioConf

type MinioConf struct {
	Host      string
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
}

// LoadMinioConfig 加载 Minio 配置
func LoadMinioConfig(configPath string) {
	logs.Info("----> load minio config from: %s", configPath)
	BConfig, err := config.NewConfig("ini", configPath)
	if err != nil {
		fmt.Println("config init error:", err)
		return
	}
	MinioConfig.Host = BConfig.String("minio::weattech_res_host")
	MinioConfig.Endpoint = BConfig.String("minio::minio_endpoint")
	MinioConfig.AccessKey = BConfig.String("minio::minio_access_key")
	MinioConfig.SecretKey = BConfig.String("minio::minio_secret_key")
	MinioConfig.Bucket = BConfig.String("minio::minio_bucket")

	logs.Info("----> load minio config success.")
}
