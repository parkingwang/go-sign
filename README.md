# Go Sign

Http API 签名验证工具包，提供对API请求的签名生成、签名校验等工具类。

# 0x00 安装

```
$ go get -u github.com/parkingwang/go-sign
```

OR

```
dep ensure --add github.com/parkingwang/go-sign
```

# 0x01 生成签名信息

`sign.GoSigner`工具类，默认支持两种签名生成算法：

1. MD5: 见 `NewGoSignerMd5()` 函数
2. Sha1 + Hmac: 见：`NewGoSignerHmac()`函数

如果需要使用其它签名生成算法，使用 `NewGoSigner(FUNC)` 指定实现签名生成算法的实现即可。

## Usage

```go
gos := NewGoSignerMd5()

// 设置签名基本参数
gos.SetAppId("9d8a121ce581499d")
gos.SetTimeStamp(1532585241)
gos.SetNonceStr("ibuaiVcKdpRxkhJA")

// 设置参与签名的其它参数
gos.AddBody("plate_number", "豫A66666")

// AppSecretKey，前后包装签名体字符串
gos.SetAppSecretWrapBody("d93047a4d6fe6111")

fmt.Println("生成签字字符串：" + gos.GetUnsignedString())
fmt.Println("输出URL字符串：" + gos.GetSignedQuery())
```

输出结果为：

> 生成签字字符串：d93047a4d6fe6111appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666&time_stamp=1532585241d93047a4d6fe6111
>
> 输出URL字符串：appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4

# 0x02 校验签名信息

`sign.GoVerifier` 工具类，用来校验签名参数的格式和时间戳。它与GoSigner一起使用，用于服务端校验API请求的签名信息。

## Usage

```go
    requestUri := "/restful/api/numbers?appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666" +
		"&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4"

	// 第一步：创建GoVerifier校验类
	verifier := NewGoVerifier()

	// 假定从RequestUri中读取校验参数
	if err := verifier.ParseQuery(requestUri); nil != err {
		t.Fatal(err)
	}

	// 或者使用verifier.ParseValues(Values)来解析。

	// 第二步：（可选）校验是否包含签名校验必要的参数
	if err := verifier.MustHasOtherFields("plate_number"); nil != err {
		t.Fatal(err)
	}

	// 第三步：检查时间戳是否超时。

	// 时间戳超时：5分钟
	verifier.SetTimeout(time.Minute * 5)
	if err := verifier.CheckTimeStamp(); nil != err {
		t.Fatal(err)
	}

	// 第四步: 创建GoSigner来重现客户端的签名信息
	signer := NewGoSignerMd5()

	// 第五步：从GoVerifier中读取所有请求参数
	signer.SetBody(verifier.GetBodyWithoutSign())

	// 第六步：从数据库读取AppID对应的SecretKey
	// appid := verifier.GetAppId()
	secretKey := "d93047a4d6fe6111"

	// 使用同样的WrapBody方式
	signer.SetAppSecretWrapBody(secretKey)

	// 服务端根据客户端参数生成签名
	sign := signer.GetSignature()

    // 最后，比较服务端生成的签名信息，与客户端提供的签名是否一致即可。
	if verifier.MustString("sign") != sign {
		t.Fatal("校验失败")
	}

```

# 注意事项

1. 请求参数是多值(Slice)的，在进入签名参数体时，会将它们以英文逗号JOIN成单独字符串再进行签名。
2. 获取Int64类型参数时，如果参数对应的值不存在、不可转换成Int64类型，会返回默认值0；


# License

```
   Copyright 2018 Xi'An iRain IoT Technology Service Co.,Ltd

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

```