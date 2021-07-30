package monitor

import (
	"huangdong2012/filecoin-monitor/metric"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/trace"
	"sync"
)

var (
	once = &sync.Once{}
)

func Init(baseOpt *model.BaseOptions, traceOpt *model.TraceOptions, metricOpt *model.MetricOptions) {
	once.Do(func() {
		model.InitBaseOptions(baseOpt)
		if traceOpt != nil {
			trace.Init(baseOpt, traceOpt)
		}
		if metricOpt != nil {
			metric.Init(baseOpt, metricOpt)
		}
	})
}
