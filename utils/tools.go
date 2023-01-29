package utils

import (
	"crypto/md5"
	"fmt"
)

// GetMd5Str 根据传入字符串获取MD5加密后的长度32位字符串
func GetMd5Str(str string) string {
	data := []byte(str)
	md5Ret := md5.Sum(data)
	return fmt.Sprintf("%x", md5Ret)
}
