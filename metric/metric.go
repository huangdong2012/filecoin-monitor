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
	if len(metricOpt.Exchange) == 0 || len(metricOpt.RouteKey) == 0 {
		panic("trace exchange or route-key invalid")
	}
	if metricOpt.PushInterval < defaultInterval {
		metricOpt.PushInterval = defaultInterval
	}
	{
		model.SetBaseOptions(baseOpt)
		options = metricOpt
		metrics.Init(wrapperGather.inner)
		initRabbit()
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
		item := &model.Metric{
			Name:   *mf.Name,
			Desc:   *mf.Help,
			Labels: parseLabels(m.Label),
			Time:   time.Now().Unix(),
		}
		switch *mf.Type {
		case dto.MetricType_COUNTER:
			item.Value = *m.Counter.Value
		case dto.MetricType_GAUGE:
			item.Value = *m.Gauge.Value
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
		out[*l.Name] = *l.Value
	}
	return out
}
