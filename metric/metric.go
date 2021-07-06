package metric

import (
	dto "github.com/prometheus/client_model/go"
	"time"
)

const (
	defaultInterval = time.Second * 10
)

var (
	options *Options
)

func Init(mq string, opt *Options) {
	if len(opt.Exchange) == 0 || len(opt.RouteKey) == 0 {
		panic("trace exchange or route-key invalid")
	}
	if opt.PushInterval < defaultInterval {
		opt.PushInterval = defaultInterval
	}
	{
		options = opt
		gatherHandler = exp.export
		initRabbit(mq)
	}

	//默认启用push-gateway主动上报的方式,如果配置了gather(prometheus主动收集)则停止主动上报
	go exp.start()
}

func Push() {
	exp.push()
}

func parseMetrics(mf *dto.MetricFamily) []*Metric {
	if mf == nil {
		return nil
	}

	out := make([]*Metric, 0, 0)
	for _, m := range mf.Metric {
		item := &Metric{
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
