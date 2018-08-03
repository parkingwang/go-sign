package sign

import (
	"fmt"
	"net/url"
	"testing"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestGoSignMd5(t *testing.T) {
	signer := NewGoSignerMd5()
	signer.SetAppId("9d8a121ce581499d")
	signer.SetTimeStamp(1532585241)
	signer.SetNonceStr("ibuaiVcKdpRxkhJA")
	signer.AddBody("plate_number", "豫A66666")
	signer.SetAppSecretWrapBody("d93047a4d6fe6111")
	fmt.Println("生成签字字符串：" + signer.GetSignBodyString())
	fmt.Println("输出URL字符串：" + signer.GetSignedQuery())
	if "appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666"+
		"&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4" != signer.GetSignedQuery() {
		t.Fatal("Md5校验失败")
	}
}

func TestGoSigner_AddBody(t *testing.T) {

	body := make(url.Values)
	body["username"] = []string{"yoojia"}
	body["tags"] = []string{"github", "gopher", "javaer"}

	signer := NewGoSignerHmac()
	signer.SetAppSecret("PASS0123")
	signer.SetTimeStamp(1234567890)
	signer.SetAppId("yoojia001")
	signer.SetNonceStr("WAHAHAH")
	for k, v := range body {
		signer.AddBodies(k, v)
	}

	body.Add(KeyNameTimeStamp, "1234567890")
	body.Add(KeyNameAppId, "yoojia001")
	body.Add(KeyNameNonceStr, "WAHAHAH")

	fmt.Println("生成签字字符串：" + signer.GetSignBodyString())
	fmt.Println("输出URL字符串：" + signer.GetSignedQuery())

	verifier := NewGoVerifier()
	verifier.ParseValues(body)

	resigner := NewGoSignerHmac()
	resigner.SetAppSecret("PASS0123")
	resigner.SetBody(verifier.GetBodyWithoutSign())

	fmt.Println("重新生成签字字符串：" + resigner.GetSignBodyString())
	fmt.Println("重新输出URL字符串：" + resigner.GetSignedQuery())
}
