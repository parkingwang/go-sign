package sign

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

var (
	fieldNameTimestamp = "time_stamp"
	fieldNameNonceStr  = "nonce_str"
	fieldNameAppId     = "appid"
	fieldNameSign      = "sign"
)

func SetFieldNameTimestamp(name string) {
	fieldNameTimestamp = name
}

func SetFieldNameNonceStr(name string) {
	fieldNameNonceStr = name
}

func SetFieldNameAppId(name string) {
	fieldNameAppId = name
}

func SetFieldNameSign(name string) {
	fieldNameSign = name
}
