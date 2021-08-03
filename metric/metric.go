package metric

import (
	dto "github.com/prometheus/client_model/go"
	"github.com/sirupsen/logrus"
	"huangdong2012/filecoin-monitor/metric/metrics"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/utils"
	"time"
)

const (
	defaultInterval = time.Second * 10
)

var (
	options *model.MetricOptions
	logger  *logrus.Entry
)

func Init(baseOpt *model.BaseOptions, metricOpt *model.MetricOptions) {
	if metricOpt.PushInterval < defaultInterval {
		metricOpt.PushInterval = defaultInterval
	}
	{
		model.InitBaseOptions(baseOpt)
		options = metricOpt
		log, err := utils.CreateLog(baseOpt.LogDir, baseOpt.LogMetricName, logrus.TraceLevel, true)
		if err != nil {
			panic(err)
		}
		fields := logrus.Fields{"room-id": baseOpt.RoomID}
		if len(baseOpt.MinerID) > 0 {
			fields["miner-id"] = baseOpt.MinerID
		}
		logger = log.WithFields(fields)

		metrics.Init(wrapperGather.inner)
	}

	// 如果配置了gather(prometheus主动收集)则停止主动上报
	go exp.start()
	logger.Info("monitor-metric starting...")
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
