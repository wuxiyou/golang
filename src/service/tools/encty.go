package tools

import (
	"sort"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)
// https://studygolang.com/articles/3447
func DoMd5(m map[string]string) string {

	keys := make([]string, 0)
	for key, val := range m {
		if len(val) <= 0 {
			delete(m, key)
			continue
		}
		if key == "sign" || key == "Sign" {
			delete(m, key)
			continue
		}

		keys = append(keys, key)
	}
	sort.Strings(keys)
	var str string
	for _, value := range keys {
		str += value +"="+ m[value] + "&"
	}
	LogInfo(str)
	s := strings.TrimRight(str, "&")

	md5Inst := md5.New()
	md5Inst.Write([]byte(s))
	result := md5Inst.Sum([]byte(""))
	return hex.EncodeToString(result)
}

// 排序
func sortInfo(m map[string]string) map[string]string {
	 var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	newMap := make(map[string]string)
	for _,k := range keys {
		newMap[k] = m[k]
	}
	return newMap
}

func connectInfo(m map[string]string) string {
	// 如果为空值则删除
	for k, v := range m {
		if len(v) <= 0 {
			delete(m, m[k])
		}
	}
	newM := sortInfo(m)
	var str string
	for key, val := range newM {
		str += key + "=" + val + "&"
	}
	return strings.TrimRight(str, "&")
}

func GeneratRandomNumber(start, end, count int) []int {
	if end < start || (end - start) < count {
		return nil
	}

	nums := make([]int, 0)

	// 随机数生成器，加入时间
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(nums) < count {
		// 随机数
		num := r.Intn((end - start)) + start

		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}