package utils

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"sync"
)

// FileIsExist 检查文件是否存在
func FileIsExist(path string) bool {
	lock := sync.RWMutex{}
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	return true
}

// FileSize 获取文件大小，单位字节
func FileSize(path string) (int64, error) {
	lock := sync.RWMutex{}
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	ff, err := f.Stat()
	if err != nil {
		return 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logs.Error(err.Error())
		}
	}(f)
	return ff.Size(), nil
}

// PathIsExist 检查目录是否存在
func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// CreatePathIfNotExist 检查目录是否存在，如果不存在则创建
func CreatePathIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// 如果不存在
	if os.IsNotExist(err) {
		// 创建目录
		err = os.MkdirAll(path, 0777)
		if err != nil {
			logs.Info("创建目录失败，path: %s, error: %v", path, err)
			return false, err
		}
	}
	return true, nil
}

// CopyFile 复制文件
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			logs.Info(err)
		}
	}(source)

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func(destination *os.File) {
		err := destination.Close()
		if err != nil {
			logs.Info(err)
		}
	}(destination)
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// DeleteFile 删除文件
func DeleteFile(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	isDir := info.IsDir()
	if isDir {
		return errors.New(filePath + " - 是目录，不是文件")
	}
	err = os.Remove(filePath)
	if err != nil {
		return errors.New("删除 - " + filePath + " 失败，error: " + err.Error())
	}
	return nil
	//exist := FileIsExist(filePath)
	//if !exist {
	//	return errors.New(filePath + " - 文件不存在")
	//}
}
