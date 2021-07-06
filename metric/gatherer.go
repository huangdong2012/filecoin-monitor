package metric

import (
	"grandhelmsman/filecoin-monitor/metric/metrics"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

type GatherHandler func(ms []*dto.MetricFamily)

var (
	gatherHandler GatherHandler = nil
)

func InitGatherer(reg *prometheus.Registry) prometheus.Gatherer {
	gather := &WrapperGatherer{outer: reg}
	if gather.outer == nil {
		gather.outer = prometheus.DefaultRegisterer.(*prometheus.Registry)
	}
	{
		//如果使用prometheus主动收集,就停止使用push-gateway上报了
		exp.stop()
	}

	return gather
}

type WrapperGatherer struct {
	outer *prometheus.Registry
}

func (g *WrapperGatherer) Gather() ([]*dto.MetricFamily, error) {
	var (
		err     error
		msInner []*dto.MetricFamily
		msOuter []*dto.MetricFamily
		out     = make([]*dto.MetricFamily, 0, 0)
	)
	if msInner, err = metrics.InnerMetrics(); err != nil {
		return nil, err
	}
	out = append(out, msInner...)
	if g.outer != nil {
		if msOuter, err = g.outer.Gather(); err != nil {
			return nil, err
		}
		out = append(out, msOuter...)
	}
	if gatherHandler != nil { //prometheus主动收集的时候，同时上报msInner到mq
		gatherHandler(msInner)
	}

	return out, nil
}
