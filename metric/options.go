package metric

import (
	"time"
)

type Options struct {
	Exchange string
	RouteKey string

	PushUrl      string        //push-gateway url
	PushJob      string        //push-gateway job name
	PushInterval time.Duration //上报metric间隔
}

type Metric struct {
	Name   string
	Desc   string
	Value  float64
	Time   int64
	Labels map[string]string
}
