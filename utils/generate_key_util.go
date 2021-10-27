package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
)

// GenerateKeyWithHmacMd5
// 使用 hmac 的算法对数据进行加密，返回加密后的字符串
func GenerateKeyWithHmacMd5(publicKey string, privateKey string) string {
	// 接收一个算法和秘钥-私钥
	hmacTool := hmac.New(md5.New, []byte(privateKey))
	// 对公钥进行加密
	hmacTool.Write([]byte(publicKey))
	var result = base64.StdEncoding.EncodeToString(hmacTool.Sum(nil))
	return result
}

// GenerateKeyWithHmacSha1
// 使用 hmac 的算法对数据进行加密，返回加密后的字符串
func GenerateKeyWithHmacSha1(publicKey string, privateKey string) string {
	// 接收一个算法和秘钥-私钥
	hmacTool := hmac.New(sha1.New, []byte(privateKey))
	// 对公钥进行加密
	hmacTool.Write([]byte(publicKey))
	var result = base64.StdEncoding.EncodeToString(hmacTool.Sum(nil))
	return result
}
