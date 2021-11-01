package utils

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/minio/minio-go/v6"
	"github.com/cuilan/go-common/load"
	"regexp"
	"strings"
)

// UploadToMinio 上传文件至 minio
// saveName: 文件路径名称，要求至少包含3个/
// localFilePath: 文件本地路径
func UploadToMinio(saveName, localFilePath string) error {
	// 上传路径应至少包含三级目录
	// ([^\\:*<>|"?\r\n/]+|\/[^\\:*<>|"?\r\n/]+)(\/[^\\:*<>|"?\r\n/]+){2}
	saveNameReg := "([^\\\\:*<>|\"?\\r\\n/]+|\\/[^\\\\:*<>|\"?\\r\\n/]+)(\\/[^\\\\:*<>|\"?\\r\\n/]+){2}"
	nameRegxp := regexp.MustCompile(saveNameReg)
	if !nameRegxp.MatchString(saveName) {
		return errors.New("saveName 至少应包含三级目录")
	}

	// Initialize minio client object.
	minioClient, err := minio.New(load.MinioConfig.Endpoint, load.MinioConfig.AccessKey, load.MinioConfig.SecretKey, false)
	if err != nil {
		logs.Error("Initialize minio client object error: %v", err.Error())
		return err
	}

	// 检查存储桶是否已存在
	exists, errBucketExists := minioClient.BucketExists(load.MinioConfig.Bucket)
	if errBucketExists != nil || !exists {
		logs.Error("Bucket is not exist, or error: %v", err)
		return errBucketExists
	}

	// 文件类型
	var contentType string
	split := strings.Split(saveName, ".")
	suffix := split[1]
	if suffix == "png" {
		contentType = "image/png"
	} else if suffix == "jpg" {
		contentType = "image/jpeg"
	}

	// 上传文件
	n, errPutObj := minioClient.FPutObject(load.MinioConfig.Bucket, saveName, localFilePath, minio.PutObjectOptions{ContentType: contentType})
	if errPutObj != nil {
		logs.Error("Put Object error, fileName: %s, error: %v", saveName, errPutObj)
	}
	logs.Info("Successfully uploaded %s of size %d", saveName, n)
	return nil
}
