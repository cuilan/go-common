package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ErrorNotice 发送异常通知
// noticeUrl: http://10.110.1.141:8080/api/notice/fsToGroup
func ErrorNotice(msg, noticeUrl string) {
	if len(noticeUrl) == 0 {
		logs.Error("noticeUrl错误.")
		return
	}
	//这里添加post的body内容
	params := url.Values{}
	params.Add("msg", msg)

	//把post表单发送给目标服务器
	res, err := http.PostForm(noticeUrl, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logs.Error("异常通知失败", err.Error())
		}
	}(res.Body)
	result, _ := ioutil.ReadAll(res.Body)
	fmt.Println("post send success", string(result))
	logs.Info("异常通知", string(result))
}
