package monitor

import (
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace"
	"sync"
)

var (
	once = &sync.Once{}
)

func Init(baseOpt *model.BaseOptions, traceOpt *model.TraceOptions, metricOpt *model.MetricOptions) {
	once.Do(func() {
		if baseOpt == nil || len(baseOpt.Node) == 0 {
			panic("base options invalid")
		}

		model.InitBaseOptions(baseOpt)
		if traceOpt != nil {
			trace.Init(baseOpt, traceOpt)
		}
		if metricOpt != nil {
			metric.Init(baseOpt, metricOpt)
		}
	})
}
