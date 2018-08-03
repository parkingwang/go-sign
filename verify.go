package sign

import (
	"time"
	"net/url"
	"errors"
	"fmt"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type GoVerify struct {
	body map[string]interface{}

	timeout time.Duration // 签名过期时间
}

func NewGoVerify() *GoVerify {
	return &GoVerify{
		body:    make(map[string]interface{}),
		timeout: time.Minute * 5,
	}
}

// SetTimeout 设置签名校验过期时间
func (slf *GoVerify) SetTimeout(timeout time.Duration) *GoVerify {
	slf.timeout = timeout
	return slf
}

// Parse 将参数字符串解析成参数列表
func (slf *GoVerify) ParseQuery(requestUri string) error {
	requestQuery := ""
	idx := strings.Index(requestUri, "?")
	if idx > 0 {
		requestQuery = requestUri[idx+1:]
	}
	query, err := url.ParseQuery(requestQuery)
	if nil != err {
		return err
	}
	// 只处理只包含一个值的参数
	for k, v := range query {
		if len(v) == 1 {
			slf.body[k] = v[0]
		} else {
			return errors.New(fmt.Sprintf("PARAM_MULTI_VALUES:<%s>", k))
		}
	}
	return nil
}

// MustString 获取字符串值
func (slf *GoVerify) MustString(key string) string {
	return fmt.Sprintf("%v", slf.body[key])
}

// MustInt64 获取Int64值
func (slf *GoVerify) MustInt64(key string) int64 {
	return convertToInt64(slf.body[key])
}

// MustHasFields 必须包含指定的字段参数
func (slf *GoVerify) MustHasFields(keys ...string) error {
	for _, key := range keys {
		if _, hit := slf.body[key]; !hit {
			return errors.New(fmt.Sprintf("PARAM_MISSED:<%s>", key))
		}
	}
	return nil
}

// MustHasFields 必须包含除特定的[time_stamp, nonce_str, sign, appid]等之外的指定的字段参数
func (slf *GoVerify) MustHasOtherFields(keys ...string) error {
	fields := []string{fieldNameTimestamp, fieldNameNonceStr, fieldNameSign, fieldNameAppId}
	fields = append(fields, keys...)
	return slf.MustHasFields(fields...)
}

// 检查时间戳有效期
func (slf *GoVerify) CheckTimeStamp() error {
	timestamp := convertToInt64(slf.body[fieldNameTimestamp])
	thatTime := time.Unix(timestamp, 0)
	if time.Now().Sub(thatTime) > slf.timeout {
		return errors.New(fmt.Sprintf("TIMESTAMP_TIMEOUT<%d>", timestamp))
	}
	return nil
}

// GetBodyWithoutSign 获取所有参数体。其中不包含sign 字段
func (slf *GoVerify) GetBodyWithoutSign() map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range slf.body {
		if k != fieldNameSign {
			out[k] = v
		}
	}
	return out
}
