package test

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/cuilan/go-common/load"
	"github.com/cuilan/go-common/utils"
	"regexp"
	"testing"
	"time"
)

// common_util.go

// file_util.go

func TestFileIsExist(t *testing.T) {
	exist := utils.FileIsExist("/Users/zhangyan/1.txt")
	fmt.Println(exist)
}

func TestFileSize(t *testing.T) {
	size, err := utils.FileSize("/Users/zhangyan/1.txt")
	fmt.Println(size)
	fmt.Println(err)
}

func TestPathIsExist(t *testing.T) {
	exist := utils.PathIsExist("/Users/zhangyan/")
	fmt.Println(exist)
}

func TestCreatePathIfNotExist(t *testing.T) {
	exist, err := utils.CreatePathIfNotExist("/tmp/zhangyan/test")
	fmt.Println(exist, " - ", err)
}

func TestCopyFile(t *testing.T) {
	b, err := utils.CopyFile("/Users/zhangyan/Downloads/1.png", "/Users/zhangyan/1.png")
	fmt.Println(b)
	fmt.Println(err)
}

func TestDeleteFile(t *testing.T) {
	err := utils.DeleteFile("/Users/zhangyan/tmp/1.txt")
	if err != nil {
		fmt.Println("删除失败", err)
	} else {
		fmt.Println("删除成功")
	}
}

// generate_key_util.go

func TestGenerateKeyWithHmacMd5(t *testing.T) {
	md5 := utils.GenerateKeyWithHmacMd5("abc", "123")
	fmt.Println(md5)
}

func TestGenerateKeyWithHmacSha1(t *testing.T) {
	sha1 := utils.GenerateKeyWithHmacSha1("abc", "123")
	fmt.Println(sha1)
}

// goftp_util.go

func TestGetFtpClient(t *testing.T) {
	ftp, _ := utils.GetFtpClient("ftp.ncep.noaa.gov", "anonymous", "anonymous", 21, false)
	fmt.Println(ftp.List("/pub/data/nccf/com/gfs/prod/"))
}

// minio_util.go

func TestRes(t *testing.T) {
	// ([^\\:*<>|"?\r\n/]+|\/[^\\:*<>|"?\r\n/]+)(\/[^\\:*<>|"?\r\n/]+){2}
	saveNameReg := "([^\\\\:*<>|\"?\\r\\n/]+|\\/[^\\\\:*<>|\"?\\r\\n/]+)(\\/[^\\\\:*<>|\"?\\r\\n/]+){2}"
	nameRegxp := regexp.MustCompile(saveNameReg)
	if !nameRegxp.MatchString("a/b/c/111.png") {
		fmt.Println("不匹配")
	} else {
		fmt.Println("匹配")
	}
}

func TestUploadToMinio(t *testing.T) {
	load.LoadMinioConfig("../conf/minio.conf")
	err := utils.UploadToMinio("/2021/10/25/1.jpg", "/Users/zhangyan/tmp/1.jpg")
	if err != nil {
		logs.Error("上传失败", err)
	} else {
		logs.Info("上传成功")
	}
}

// notice.go

func TestErrorNotice(t *testing.T) {
	utils.ErrorNotice("go-common 测试消息", "http://10.110.1.141:8080/api/notice/fsToGroup")
}

// time_util.go

func TestTimeUtil(t *testing.T) {
	fmt.Println(utils.Now2SecondsTimestamp())
	fmt.Println(utils.Now2MilliTimestamp())

	format := time.Unix(1595061723, 0)
	fmt.Println(format)
}
