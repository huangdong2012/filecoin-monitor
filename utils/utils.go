package utils

import maddr "github.com/micro/go-micro/v2/util/addr"

var (
	ipAddr = "127.0.0.1"
)

func IpAddr() string {
	if len(ipAddr) > 0 && ipAddr != "127.0.0.1" {
		return ipAddr
	}

	addr, err := maddr.Extract("0.0.0.0")
	if err == nil {
		ipAddr = addr
	}
	return ipAddr
}

func GetKeys(kvs map[string]string) []string {
	keys := make([]string, 0, 0)
	for k, _ := range kvs {
		keys = append(keys, k)
	}
	return keys
}

func GetValues(kvs map[string]string) []string {
	values := make([]string, 0, 0)
	for _, v := range kvs {
		values = append(values, v)
	}
	return values
}
