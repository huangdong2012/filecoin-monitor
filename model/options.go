package model

import (
	"grandhelmsman/filecoin-monitor/utils"
	"sync"
	"time"
)

var (
	base *BaseOptions
	once = &sync.Once{}
)

func InitBaseOptions(baseOpt *BaseOptions) {
	once.Do(func() {
		base = baseOpt
		utils.InitLog(baseOpt.LogErr, baseOpt.LogInfo)
	})
}

func GetBaseOptions() *BaseOptions {
	return base
}

type BaseOptions struct {
	RoomID  int64  //机房ID
	Role    Role   //软件包类型
	MinerID string //如:t01000

	LogErr  func(error)
	LogInfo func(string)
}

type TraceOptions struct {
	ExportSpan func(span *Span)
}

type MetricOptions struct {
	PushUrl      string        //push-gateway url
	PushJob      string        //push-gateway job name
	PushInterval time.Duration //上报metric间隔

	ExportMetric func(metrics []*Metric)
}
