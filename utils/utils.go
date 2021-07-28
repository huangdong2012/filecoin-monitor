package utils

import (
	"encoding/json"
	maddr "github.com/micro/go-micro/v2/util/addr"
	"os"
	"time"
)

var (
	ipAddr = "127.0.0.1"
)

func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

func IpAddr() string {
	if len(ipAddr) > 0 && ipAddr != "127.0.0.1" {
		return ipAddr
	}

	if addr, err := maddr.Extract("0.0.0.0"); err == nil {
		ipAddr = addr
	}
	return ipAddr
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetMapKeys(kvs map[string]string) []string {
	keys := make([]string, 0, 0)
	for k, _ := range kvs {
		keys = append(keys, k)
	}
	return keys
}

func GetMapValues(kvs map[string]string) []string {
	values := make([]string, 0, 0)
	for _, v := range kvs {
		values = append(values, v)
	}
	return values
}

func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
	}
	return false
}

func StructToMap(s interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	if s == nil {
		return out
	}
	if data, err := json.Marshal(s); err == nil {
		_ = json.Unmarshal(data, &out)
	}
	return out
}
