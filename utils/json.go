package utils

import "encoding/json"

func ToJson(obj interface{}) (string, error) {
	data, err := json.MarshalIndent(obj, "", "	")
	return string(data), err
}

func ToJsonWithoutError(obj interface{}) string {
	data, _ := json.MarshalIndent(obj, "", "	")
	return string(data)
}
