package metric

import (
	dto "github.com/prometheus/client_model/go"
	"grandhelmsman/filecoin-monitor/metric/metrics"
	"grandhelmsman/filecoin-monitor/model"
	"time"
)

const (
	defaultInterval = time.Second * 10
)

var (
	options *model.MetricOptions
)

func Init(baseOpt *model.BaseOptions, metricOpt *model.MetricOptions) {
	if metricOpt.PushInterval < defaultInterval {
		metricOpt.PushInterval = defaultInterval
	}
	{
		model.InitBaseOptions(baseOpt)
		options = metricOpt
		metrics.Init(wrapperGather.inner)
	}

	//默认启用push-gateway主动上报的方式,如果配置了gather(prometheus主动收集)则停止主动上报
	go exp.start()
}

func Push() {
	exp.pushAll()
}

func PushScope(s *Scope) {
	if s != nil {
		s.Push()
	}
}

func parseMetrics(mf *dto.MetricFamily) []*model.Metric {
	if mf == nil {
		return nil
	}

	out := make([]*model.Metric, 0, 0)
	for _, m := range mf.Metric {
		if m == nil {
			continue
		}
		item := &model.Metric{
			Name:   mf.GetName(),
			Desc:   mf.GetHelp(),
			Labels: parseLabels(m.Label),
			Time:   time.Now().Unix(),
		}
		switch mf.GetType() {
		case dto.MetricType_COUNTER:
			item.Value = m.Counter.GetValue()
		case dto.MetricType_GAUGE:
			item.Value = m.Gauge.GetValue()
		default:
			continue
		}

		out = append(out, item)
	}
	return out
}

func parseLabels(ls []*dto.LabelPair) map[string]string {
	out := make(map[string]string)
	for _, l := range ls {
		if l != nil {
			out[l.GetName()] = l.GetValue()
		}
	}
	return out
}
