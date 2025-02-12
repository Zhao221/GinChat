package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// MakePassword 加密
func MakePassword(plainPassword string, salt string) string {
	return Md5Encode(plainPassword + salt)
}

// ValidPassword 解密
func ValidPassword(plainPassword string, salt string, password string) bool {
	return Md5Encode(plainPassword+salt) == password
}
