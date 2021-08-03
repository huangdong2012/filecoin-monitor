package model

import (
	"sync"
	"time"
)

var (
	base *BaseOptions
	once = &sync.Once{}
)

func InitBaseOptions(baseOpt *BaseOptions) {
	once.Do(func() {
		if baseOpt == nil {
			panic("base options invalid")
		}
		if baseOpt.PackageKind == PackageKind_Miner && len(baseOpt.MinerID) == 0 {
			panic("miner-id invalid")
		}
		if len(baseOpt.LogTraceName) == 0 {
			baseOpt.LogTraceName = "monitor-trace"
		}
		if len(baseOpt.LogMetricName) == 0 {
			baseOpt.LogMetricName = "monitor-metric"
		}
		base = baseOpt
	})
}

func GetBaseOptions() *BaseOptions {
	return base
}

type BaseOptions struct {
	RoomID      int64       //机房ID
	HostNo      string      //主机编号
	MinerID     string      //如:t01000
	PackageKind PackageKind //软件包类型

	LogDir        string //日志文件夹名称
	LogTraceName  string //trace日志文件名
	LogMetricName string //metric日志文件名
}

type TraceOptions struct {
	ExportAll      bool //true: 导出所有的span  false: 只导出monitor定义的span
	ExportSpan     func(span *Span)

	SpanLogDir  string
	SpanLogName string
}

type MetricOptions struct {
	PushUrl      string        //push-gateway url
	PushJob      string        //push-gateway job name
	PushInterval time.Duration //上报metric间隔

	ExportMetric   func(metrics []*Metric)
}
