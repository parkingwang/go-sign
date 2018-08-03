package sign

import (
	"fmt"
	"testing"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func TestGoVerify_ParseQuery(t *testing.T) {
	requestUri := "/restful/api/numbers?appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666" +
		"&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4"

	// 第一步：创建GoVerify校验类
	verifier := NewGoVerifier()

	// 假定从RequestUri中读取校验参数
	if err := verifier.ParseQuery(requestUri); nil != err {
		t.Fatal(err)
	}

	// 第二步：（可选）校验是否包含签名校验必要的参数
	if err := verifier.MustHasOtherKeys("plate_number"); nil != err {
		t.Fatal(err)
	}

	// 第三步：检查时间戳是否超时。
	//if err := verifier.CheckTimeStamp(); nil != err {
	//	t.Fatal(err)
	//}

	// 第四步，创建GoSign来重现客户端的签名信息：
	signer := NewGoSignerMd5()

	// 第五步：从GoVerify中读取所有请求参数
	signer.SetBody(verifier.GetBodyWithoutSign())

	// 第六步：从数据库读取AppID对应的SecretKey
	// appid := verifier.MustString("appid")
	secretKey := "d93047a4d6fe6111"

	// 使用同样的WrapBody方式
	signer.SetAppSecretWrapBody(secretKey)

	// 生成
	sign := signer.GetSignature()

	if verifier.MustString("sign") != sign {
		t.Fatal("校验失败")
	}

	fmt.Println(sign)
}
