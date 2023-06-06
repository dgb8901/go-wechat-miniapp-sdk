package common

import (
	"encoding/json"
	"log"
)

// ToJson 对象转json
func ToJson(md interface{}) string {
	bytes, _ := json.Marshal(md)
	return string(bytes)
}

// JsonToMap json 转 map
func JsonToMap(data string) map[string]interface{} {

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(data), &mapResult); err != nil {
		log.Print(err)
	}
	return mapResult
}

// IsBlank 判断字符串为空
func IsBlank(s string) bool {
	return s == ""
}

// IsNotBlank 判断字符串不为空
func IsNotBlank(s string) bool {
	return s != ""
}
