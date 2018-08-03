package sign

import (
	"sort"
	"strings"
	"fmt"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
// API签名
//

// 签名加密函数
type CryptoFunc func(secretKey string, args string) []byte

type GoSign struct {
	body map[string]interface{}

	bodyPrefix string // 参数体前缀
	bodySuffix string // 参数体后缀
	splitChar  string // 前缀、后缀分隔符号

	secretKey  string // 签名密钥
	cryptoFunc CryptoFunc
}

func NewGoSign(cryptoFunc CryptoFunc) *GoSign {
	return &GoSign{
		body:       make(map[string]interface{}),
		bodyPrefix: "",
		bodySuffix: "",
		splitChar:  "",
		cryptoFunc: cryptoFunc,
	}
}

func NewGoSignMd5() *GoSign {
	return NewGoSign(Md5Sign)
}

func NewGoSignHmac() *GoSign {
	return NewGoSign(Hmac5Sign)
}

// SetBody 设置整个参数体Body对象。
func (slf *GoSign) SetBody(body map[string]interface{}) {
	slf.body = body
}

// AddBody 添加签名体字段和值
func (slf *GoSign) AddBody(key string, value interface{}) *GoSign {
	slf.body[key] = value
	return slf
}

// SetTimeStamp 设置时间戳参数
func (slf *GoSign) SetTimeStamp(ts int64) *GoSign {
	return slf.AddBody(fieldNameTimestamp, ts)
}

// SetNonceStr 设置随机字符串参数
func (slf *GoSign) SetNonceStr(nonce string) *GoSign {
	return slf.AddBody(fieldNameNonceStr, nonce)
}

// SetAppId 设置AppId参数
func (slf *GoSign) SetAppId(appId string) *GoSign {
	return slf.AddBody(fieldNameAppId, appId)
}

// RandNonceStr 自动生成16位随机字符串参数
func (slf *GoSign) RandNonceStr() *GoSign {
	return slf.SetNonceStr(RandString(16))
}

// SetSignBodyPrefix 设置签名字符串的前缀字符串
func (slf *GoSign) SetSignBodyPrefix(prefix string) *GoSign {
	slf.bodyPrefix = prefix
	return slf
}

// SetSignBodySuffix 设置签名字符串的后缀字符串
func (slf *GoSign) SetSignBodySuffix(suffix string) *GoSign {
	slf.bodySuffix = suffix
	return slf
}

// SetSplitChar设置前缀、后缀与签名体之间的分隔符号。默认为空字符串
func (slf *GoSign) SetSplitChar(split string) *GoSign {
	slf.splitChar = split
	return slf
}

// SetAppSecret 设置签名密钥
func (slf *GoSign) SetAppSecret(appSecret string) *GoSign {
	slf.secretKey = appSecret
	return slf
}

// SetAppSecretWrapBody 在签名参数体的首部和尾部，拼接AppSecret字符串。
func (slf *GoSign) SetAppSecretWrapBody(appSecret string) *GoSign {
	slf.SetSignBodyPrefix(appSecret)
	slf.SetSignBodySuffix(appSecret)
	return slf.SetAppSecret(appSecret)
}

// GetSignBodyString 获取用于签名的原始字符串
func (slf *GoSign) GetSignBodyString() string {
	return slf.bodyPrefix + slf.splitChar + slf.getSortedBodyString() + slf.splitChar + slf.bodySuffix
}

// GetSignedQuery 获取带签名参数的字符串
func (slf *GoSign) GetSignedQuery() string {
	body := slf.getSortedBodyString()
	sign := slf.GetSignature()
	return body + "&" + fieldNameSign + "=" + sign
}

// GetSignature 获取签名字符串
func (slf *GoSign) GetSignature() string {
	sign := fmt.Sprintf("%x", slf.cryptoFunc(slf.secretKey, slf.GetSignBodyString()))
	return sign
}

func (slf *GoSign) getSortedBodyString() string {
	return SortKVPairs(slf.body)
}

////

// SortKVPairs 将Map的键值对，按字典顺序拼接成字符串
func SortKVPairs(m map[string]interface{}) string {
	size := len(m)
	if size == 0 {
		return ""
	}
	keys := make([]string, size)
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	pairs := make([]string, size)
	for i, key := range keys {
		pairs[i] = fmt.Sprintf("%s=%v", key, m[key])
	}
	return strings.Join(pairs, "&")
}
