package sign

import (
	"fmt"
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestGoSignMd5(t *testing.T) {
	gos := NewGoSignerMd5()
	gos.SetAppId("9d8a121ce581499d")
	gos.SetTimeStamp(1532585241)
	gos.SetNonceStr("ibuaiVcKdpRxkhJA")
	gos.AddBody("plate_number", "豫A66666")
	gos.SetAppSecretWrapBody("d93047a4d6fe6111")
	fmt.Println("生成签字字符串：" + gos.GetSignBodyString())
	fmt.Println("输出URL字符串：" + gos.GetSignedQuery())
	if "appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666"+
		"&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4" != gos.GetSignedQuery() {
		t.Fatal("Md5校验失败")
	}
}
