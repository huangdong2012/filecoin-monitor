package monitor

import (
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/trace"
	"grandhelmsman/filecoin-monitor/utils"
	"sync"
)

var (
	once = &sync.Once{}
)

type Options struct {
	MQUrl         string //rabbit or kafka
	TraceOptions  *trace.Options
	MetricOptions *metric.Options

	LogErr  func(error)
	LogInfo func(string)
}

func Init(opt *Options) {
	once.Do(func() {
		if opt == nil || len(opt.MQUrl) == 0 {
			panic("monitor options invalid")
		}

		if opt.TraceOptions != nil {
			trace.Init(opt.MQUrl, opt.TraceOptions)
		}

		if opt.MetricOptions != nil {
			metric.Init(opt.MQUrl, opt.MetricOptions)
		}

		utils.InitLog(opt.LogErr, opt.LogInfo)
	})
}
