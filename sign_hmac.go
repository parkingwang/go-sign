package sign

import (
	"crypto/hmac"
	"crypto/sha1"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func Hmac5Sign(secretKey, body string) []byte {
	m := hmac.New(sha1.New, []byte(secretKey))
	m.Write([]byte(body))
	return m.Sum(nil)
}
