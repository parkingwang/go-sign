package sign

import "crypto/md5"

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func Md5Sign(_, body string) []byte {
	m := md5.New()
	m.Write([]byte(body))
	return m.Sum(nil)
}
