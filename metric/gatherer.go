package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

var (
	wrapperGather = &WrapperGatherer{
		inner: prometheus.NewRegistry(),
	}
)

func InitGatherer(reg *prometheus.Registry) prometheus.Gatherer {
	if wrapperGather.outer = reg; wrapperGather.outer == nil {
		wrapperGather.outer = prometheus.DefaultRegisterer.(*prometheus.Registry)
	}
	{
		exp.stop() //如果使用prometheus主动收集,就停止使用push-gateway上报了
	}
	return wrapperGather
}

type WrapperGatherer struct {
	inner *prometheus.Registry //自定义指标的注册中心
	outer *prometheus.Registry //系统默认(open-census)指标的注册中心
}

func (g *WrapperGatherer) Gather() ([]*dto.MetricFamily, error) {
	var (
		err     error
		msInner []*dto.MetricFamily
		msOuter []*dto.MetricFamily
		out     = make([]*dto.MetricFamily, 0, 0)
	)
	if g.inner != nil {
		if msInner, err = g.inner.Gather(); err != nil {
			return nil, err
		}
		out = append(out, msInner...)
	}
	if g.outer != nil {
		if msOuter, err = g.outer.Gather(); err != nil {
			return nil, err
		}
		out = append(out, msOuter...)
	}
	if err = exp.export(g.inner); err != nil { //prometheus收集的时候，同时将自定义指标上传到mq
		return nil, err
	}
	return out, nil
}
