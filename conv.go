package sign

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

func convertToInt64(val interface{}) int64 {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case json.Number:
		i, _ := val.(json.Number).Int64()
		return i
	case int:
		return int64(val.(int))

	case int32:
		return int64(val.(int32))

	case int64:
		return val.(int64)

	case float32:
		return int64(val.(float32))

	case float64:
		return int64(val.(float64))

	default:
		sv := fmt.Sprintf("%v", val)
		if iv, err := strconv.ParseInt(sv, 10, 64); nil != err {
			return 0
		} else {
			return iv
		}
	}
}
