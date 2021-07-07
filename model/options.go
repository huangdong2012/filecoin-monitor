package model

import "time"

type BaseOptions struct {
	Role  Role
	Node  string //如: t01000
	MQUrl string //rabbit or kafka

	LogErr  func(error)
	LogInfo func(string)
}

type TraceOptions struct {
	Exchange string
	RouteKey string
}

type MetricOptions struct {
	Exchange string
	RouteKey string

	PushUrl      string        //push-gateway url
	PushJob      string        //push-gateway job name
	PushInterval time.Duration //上报metric间隔
}
