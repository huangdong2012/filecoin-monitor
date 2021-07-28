package trace

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"grandhelmsman/filecoin-monitor/utils"
	"strconv"
	"time"
)

var (
	exp = &Exporter{
		metricName:      "metric",
		metricSpanID:    "span_id",
		metricStatus:    "status",
		metricStartTime: "start_time",
		metricEndTime:   "end_time",
	}
)

type Exporter struct {
	metricName      string
	metricSpanID    string
	metricStatus    string
	metricStartTime string
	metricEndTime   string
}

// ExportSpan send to kafka
func (e *Exporter) ExportSpan(sd *trace.SpanData) {
	var (
		err  error
		span *model.Span
	)
	if !options.ExportAll && !spans.IsLocalSpan(sd) {
		return
	}
	if span, err = parseSpan(sd); err != nil {
		logger.WithField("err", err).Error("parse span data error")
		return
	} else {
		if spanLogger != nil && options.ExportSpan == nil {
			spanLogger.WithFields(utils.StructToMap(span)).Info("")
		}
	}
	if spans.MetricEnable(span.Tags) && e.pushMetricEnable(span.Tags) {
		if err = e.pushMetric(span); err != nil {
			logger.WithField("err", err).Error("span to metric error")
		}
	}
}

func (e *Exporter) pushMetricEnable(tags map[string]string) bool {
	if tagStatus, ok := tags["status"]; ok {
		if status, err := strconv.Atoi(tagStatus); err == nil {
			if status < model.TaskStatus_Finish {
				return false
			}
		}
	}
	return true
}

func (e *Exporter) pushMetric(span *model.Span) error {
	var (
		ok    bool
		name  string
		gauge *prometheus.GaugeVec
	)
	if span.Tags == nil {
		return nil
	}
	if name, ok = span.Tags[e.metricName]; !ok {
		return nil
	}

	labelValues := map[string]string{
		e.metricSpanID:    span.ID,
		e.metricStatus:    fmt.Sprintf("%v", span.Status),
		e.metricStartTime: utils.TimeFormat(time.Unix(span.StartTime, 0)),
		e.metricEndTime:   utils.TimeFormat(time.Unix(span.EndTime, 0)),
	}
	for k, v := range span.Tags {
		labelValues[k] = v
	}

	//此处使用了临时的metric(每次都重新创建)，因为: 该metric的label值每次都不同(span_id等)，
	//导致MetricFamily里面的Metrics不断递增，数据被重复收集
	gauge = e.newMetric(name, utils.GetMapKeys(span.Tags))
	gauge.With(labelValues).Set(span.Duration)
	metric.NewScope().Add(gauge).Push()

	return nil
}

func (e *Exporter) newMetric(name string, labels []string) *prometheus.GaugeVec {
	gaugeName := fmt.Sprintf("%v_%v", string(model.GetBaseOptions().PackageKind), name)
	gaugeLbs := append(labels, e.metricSpanID, e.metricStatus, e.metricStartTime, e.metricEndTime)
	return prometheus.NewGaugeVec(prometheus.GaugeOpts(e.setupMetricOptions(gaugeName)), gaugeLbs)
}

func (e *Exporter) setupMetricOptions(name string) prometheus.Opts {
	return prometheus.Opts{
		Namespace: "zdz",
		Name:      name,
		Help:      "from span",
		ConstLabels: map[string]string{
			"room_id":  strconv.FormatInt(model.GetBaseOptions().RoomID, 10),
			"instance": utils.IpAddr(),
			"host_no":  model.GetBaseOptions().HostNo,
			"miner_id": model.GetBaseOptions().MinerID,
		},
	}
}
