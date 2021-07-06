package utils

import "encoding/json"

func ToJson(obj interface{}) (string, error) {
	data, err := json.MarshalIndent(obj, "", "	")
	return string(data), err
}
