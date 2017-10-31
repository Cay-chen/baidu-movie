package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"
)

/**
	返回URL路径去除get后面带的参数
**/
func CutUrlString(s string, cuturl string) string {
	if strings.Contains(s, cuturl) {
		return strings.Split(s, cuturl)[0]
	}
	return s
}

/**
	判断文件是否存在
**/
func IsPathExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/**
三元运算符
*/
func If3(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/****
md5 返回32j机密小写
****/

func Md5R32(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
