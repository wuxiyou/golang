package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"
)
func Md5(s string) string  {
	md5Htx := md5.New()
	md5Htx.Write([]byte(s))
	return hex.EncodeToString(md5Htx.Sum(nil))
}

// 32 为
func UniqueValue() string  {
	str := make([]byte, 48)
	LogInfo(str)
	if _, err := io.ReadFull(rand.Reader, str); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(str))
}

/**
 *  判断文件是否存在， 返回bool
 */
func checkFileIsExist(filename string) bool  {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 获取当前时间
func GetCurrentTime() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t
}

