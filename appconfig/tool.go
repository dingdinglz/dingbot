package appconfig

import (
	"math/rand"
)

// 生成一个指定长度的随机字符串，内容包括26个字母和10个数字
func GenerateRandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
