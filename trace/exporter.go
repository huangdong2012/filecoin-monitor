package trace

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"grandhelmsman/filecoin-monitor/utils"
	"time"
)

func newExporter() *Exporter {
	return &Exporter{
		metricFlag:      "metric",
		metricSpanID:    "span_id",
		metricStatus:    "status",
		metricStartTime: "start_time",
		metricEndTime:   "end_time",
	}
}

type Exporter struct {
	metricFlag      string
	metricSpanID    string
	metricStatus    string
	metricStartTime string
	metricEndTime   string
}

// ExportSpan send to mq
func (e *Exporter) ExportSpan(sd *trace.SpanData) {
	var (
		err  error
		span *model.Span
		data string
	)
	if !spans.Verify(sd) {
		return
	}
	if span, err = parseSpan(sd); err != nil {
		utils.Error(fmt.Errorf("parse span data error: %v", err))
		return
	}
	if data, err = utils.ToJson(span); err != nil {
		utils.Error(fmt.Errorf("marshal span to json error: %v", err))
		return
	}
	if err = sendToRabbit([]byte(data)); err != nil {
		utils.Error(fmt.Errorf("trace exporter send to mq error: %v", err.Error()))
	}
	if err = e.pushMetric(span); err != nil {
		utils.Error(fmt.Errorf("trace exporter to metric error: %v", err.Error()))
	}
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
	if name, ok = span.Tags[e.metricFlag]; !ok {
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
	gauge = e.newMetric(name, utils.GetKeys(span.Tags))
	gauge.With(labelValues).Set(span.Duration)
	metric.NewScope().Add(gauge).Push()

	return nil
}

func (e *Exporter) newMetric(name string, labels []string) *prometheus.GaugeVec {
	gaugeName := fmt.Sprintf("%v_%v", string(model.GetBaseOptions().Role), name)
	gaugeLbs := append(labels, e.metricSpanID, e.metricStatus, e.metricStartTime, e.metricEndTime)
	return prometheus.NewGaugeVec(prometheus.GaugeOpts(e.setupMetricOptions(gaugeName)), gaugeLbs)
}

func (e *Exporter) setupMetricOptions(name string) prometheus.Opts {
	return prometheus.Opts{
		Namespace: "zdz",
		Name:      name,
		Help:      "from span",
		ConstLabels: map[string]string{
			"instance": utils.IpAddr(),
			"node":     model.GetBaseOptions().Node,
		},
	}
}
