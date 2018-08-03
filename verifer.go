package sign

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GoVerifier struct {
	*DefaultFields
	body url.Values

	timeout time.Duration // 签名过期时间
}

func NewGoVerifier() *GoVerifier {
	return &GoVerifier{
		DefaultFields: newDefaultSignFields(),
		body:          make(url.Values),
		timeout:       time.Minute * 5,
	}
}

// ParseQuery 将参数字符串解析成参数列表
func (slf *GoVerifier) ParseQuery(requestUri string) error {
	requestQuery := ""
	idx := strings.Index(requestUri, "?")
	if idx > 0 {
		requestQuery = requestUri[idx+1:]
	}
	query, err := url.ParseQuery(requestQuery)
	if nil != err {
		return err
	}
	slf.ParseValues(query)
	return nil
}

// ParseValues 将Values参数列表解析成参数Map。如果参数是多值的，则将它们以逗号Join成字符串。
func (slf *GoVerifier) ParseValues(values url.Values) {
	for key, value := range values {
		slf.body[key] = value
	}
}

// SetTimeout 设置签名校验过期时间
func (slf *GoVerifier) SetTimeout(timeout time.Duration) *GoVerifier {
	slf.timeout = timeout
	return slf
}

// MustString 获取字符串值
func (slf *GoVerifier) MustString(key string) string {
	if ss := slf.MustStrings(key); len(ss) == 0 {
		return ""
	} else {
		return ss[0]
	}
}

// MustString 获取字符串值数组
func (slf *GoVerifier) MustStrings(key string) []string {
	return slf.body[key]
}

// MustInt64 获取Int64值
func (slf *GoVerifier) MustInt64(key string) int64 {
	return convertToInt64(slf.MustString(key))
}

// MustHasKeys 必须包含指定的字段参数
func (slf *GoVerifier) MustHasKeys(keys ...string) error {
	for _, key := range keys {
		if _, hit := slf.body[key]; !hit {
			return errors.New(fmt.Sprintf("KEY_MISSED:<%s>", key))
		}
	}
	return nil
}

// MustHasKeys 必须包含除特定的[time_stamp, nonce_str, sign, appid]等之外的指定的字段参数
func (slf *GoVerifier) MustHasOtherKeys(keys ...string) error {
	fields := []string{slf.fieldNameTimestamp, slf.fieldNameNonceStr, slf.fieldNameSign, slf.fieldNameAppId}
	if len(keys) > 0 {
		fields = append(fields, keys...)
	}
	return slf.MustHasKeys(fields...)
}

// 检查时间戳有效期
func (slf *GoVerifier) CheckTimeStamp() error {
	timestamp := slf.GetTimestamp()
	thatTime := time.Unix(timestamp, 0)
	if time.Now().Sub(thatTime) > slf.timeout {
		return errors.New(fmt.Sprintf("TIMESTAMP_TIMEOUT:<%d>", timestamp))
	}
	return nil
}

func (slf *GoVerifier) GetAppId() string {
	return slf.MustString(slf.fieldNameAppId)
}

func (slf *GoVerifier) GetNonceStr() string {
	return slf.MustString(slf.fieldNameNonceStr)
}

func (slf *GoVerifier) GetSign() string {
	return slf.MustString(slf.fieldNameSign)
}

func (slf *GoVerifier) GetTimestamp() int64 {
	return slf.MustInt64(slf.fieldNameTimestamp)
}

// GetBodyWithoutSign 获取所有参数体。其中不包含sign 字段
func (slf *GoVerifier) GetBodyWithoutSign() url.Values {
	out := make(url.Values)
	for k, v := range slf.body {
		if k != slf.fieldNameSign {
			out[k] = v
		}
	}
	return out
}

func (slf *GoVerifier) GetBody() url.Values {
	out := make(url.Values)
	for k, v := range slf.body {
		out[k] = v
	}
	return out
}
