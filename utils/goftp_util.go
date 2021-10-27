package utils

import (
	"gopkg.in/dutchcoders/goftp.v1"
	"strconv"
)

// GetFtpClient 建立FTP连接
func GetFtpClient(Host, UserName, Password string, port int, IsDebug bool) (*goftp.FTP, error) {
	var err error
	var ftp *goftp.FTP
	if IsDebug {
		if ftp, err = goftp.ConnectDbg(Host + ":" + strconv.Itoa(port)); err != nil {
			panic(err)
		}
	} else {
		if ftp, err = goftp.Connect(Host + ":" + strconv.Itoa(port)); err != nil {
			panic(err)
		}
	}
	if err = ftp.Login(UserName, Password); err != nil {
		panic(err)
	}
	return ftp, nil
}
