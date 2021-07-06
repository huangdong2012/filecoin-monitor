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
