package sign

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

const (
	DefFieldNameTimeStamp = "time_stamp"
	DefFieldNameNonceStr  = "nonce_str"
	DefFieldNameAppId     = "appid"
	DefFieldNameSign      = "sign"
)

var (
	gFieldNameTimestamp = DefFieldNameTimeStamp
	gFieldNameNonceStr  = DefFieldNameNonceStr
	gFieldNameAppId     = DefFieldNameAppId
	gFieldNameSign      = DefFieldNameSign
)

func SetFieldNameTimestamp(name string) {
	gFieldNameTimestamp = name
}

func SetFieldNameNonceStr(name string) {
	gFieldNameNonceStr = name
}

func SetFieldNameAppId(name string) {
	gFieldNameAppId = name
}

func SetFieldNameSign(name string) {
	gFieldNameSign = name
}

////

type DefaultFields struct {
	fieldNameTimestamp string
	fieldNameNonceStr  string
	fieldNameAppId     string
	fieldNameSign      string
}

func newDefaultSignFields() *DefaultFields {
	return &DefaultFields{
		fieldNameTimestamp: gFieldNameTimestamp,
		fieldNameNonceStr:  gFieldNameNonceStr,
		fieldNameAppId:     gFieldNameAppId,
		fieldNameSign:      gFieldNameSign,
	}
}

func (slf *DefaultFields) SetFieldNameTimestamp(name string) {
	slf.fieldNameTimestamp = name
}

func (slf *DefaultFields) SetFieldNameNonceStr(name string) {
	slf.fieldNameNonceStr = name
}

func (slf *DefaultFields) SetFieldNameAppId(name string) {
	slf.fieldNameAppId = name
}

func (slf *DefaultFields) SetFieldNameSign(name string) {
	slf.fieldNameSign = name
}
