package sign

import (
	"math/rand"
	"time"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 生成随机字符串
//

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randomSrc = rand.NewSource(time.Now().UnixNano())

// RandString 返回指定长度的随机字符串。
// 最小长度为4，最大长度为1024
func RandString(num int) string {
	if num < 4 {
		num = 4
	}
	if num > 1024 {
		num = 1024
	}
	bytes := make([]byte, num)
	for i, cache, remain := num-1, randomSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randomSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			bytes[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(bytes)
}
