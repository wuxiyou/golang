package tools

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
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

func HttpPost(url , data string) (string, error){
	str := strings.NewReader(data)
	//fmt.Println(str)
	postReq, err := http.NewRequest("POST", url, str)
	if err != nil {
		//fmt.Println(err)
		return "", err
	}
	postReq.Header.Set("Content-Type", "application/json;encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)

	defer resp.Body.Close()

	if err != nil {
		//fmt.Println(err)
		return "", err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//fmt.Println(err)
			return "", err
		}
		//fmt.Println("POST请求:创建成功", string(body))
		//fmt.Println(err)
		return string(body), nil
	}
}

func HttpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

