package strs

import "encoding/json"

func ToJson(obj interface{}) string {
	jsonData, _ := json.Marshal(obj)
	return string(jsonData)
}

func FromJson(jsonStr string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), obj)
}
