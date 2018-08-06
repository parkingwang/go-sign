package sign

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// 提供一个修改默认字段名的扩展
//

const (
	KeyNameTimeStamp = "time_stamp"
	KeyNameNonceStr  = "nonce_str"
	KeyNameAppId     = "appid"
	KeyNameSign      = "sign"
)

var (
	gKeyNameTimestamp = KeyNameTimeStamp
	gKeyNameNonceStr  = KeyNameNonceStr
	gKeyNameAppId     = KeyNameAppId
	gKeyNameSign      = KeyNameSign
)

func SetKeyNameTimestamp(name string) {
	gKeyNameTimestamp = name
}

func SetKeyNameNonceStr(name string) {
	gKeyNameNonceStr = name
}

func SetKeyNameAppId(name string) {
	gKeyNameAppId = name
}

func SetKeyNameSign(name string) {
	gKeyNameSign = name
}

////

type DefaultKeyName struct {
	keyNameTimestamp string
	keyNameNonceStr  string
	keyNameAppId     string
	keyNameSign      string
}

func newDefaultKeyName() *DefaultKeyName {
	return &DefaultKeyName{
		keyNameTimestamp: gKeyNameTimestamp,
		keyNameNonceStr:  gKeyNameNonceStr,
		keyNameAppId:     gKeyNameAppId,
		keyNameSign:      gKeyNameSign,
	}
}

func (slf *DefaultKeyName) SetKeyNameTimestamp(name string) {
	slf.keyNameTimestamp = name
}

func (slf *DefaultKeyName) SetKeyNameNonceStr(name string) {
	slf.keyNameNonceStr = name
}

func (slf *DefaultKeyName) SetKeyNameAppId(name string) {
	slf.keyNameAppId = name
}

func (slf *DefaultKeyName) SetKeyNameSign(name string) {
	slf.keyNameSign = name
}
