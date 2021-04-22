package common

import (
	"encoding/json"
	"github.com/wonderivan/logger"
)

// 对象转json
func ToJson(md interface{}) string {
	bytes, _ := json.Marshal(md)
	return string(bytes)
}

// json 转 map
func JsonToMap(data string) map[string]interface{} {

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(data), &mapResult); err != nil {
		logger.Error(err)
	}
	return mapResult
}

// 判断字符串为空
func IsBlank(s string) bool {
	return s == ""
}

// 判断字符串不为空
func IsNotBlank(s string) bool {
	return s != ""
}
