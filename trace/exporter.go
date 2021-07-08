package trace

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/metric/metrics"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"grandhelmsman/filecoin-monitor/utils"
	"sync"
	"time"
)

func newExporter() *Exporter {
	return &Exporter{
		metricFlag:      "metric",
		metrics:         &sync.Map{},
		metricSpanID:    "span_id",
		metricStatus:    "status",
		metricStartTime: "start_time",
		metricEndTime:   "end_time",
	}
}

type Exporter struct {
	metricFlag      string
	metrics         *sync.Map
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

	gauge = e.getMetric(name, utils.GetKeys(span.Tags))
	labelValues := map[string]string{
		e.metricSpanID:    span.ID,
		e.metricStatus:    fmt.Sprintf("%v", span.Status),
		e.metricStartTime: utils.TimeFormat(time.Unix(span.StartTime,0)),
		e.metricEndTime:   utils.TimeFormat(time.Unix(span.EndTime,0)),
	}
	for k, v := range span.Tags {
		labelValues[k] = v
	}
	gauge.With(labelValues).Set(span.Duration)
	metric.Push()

	return nil
}

func (e *Exporter) getMetric(name string, labels []string) *prometheus.GaugeVec {
	var (
		ok        bool
		obj       interface{}
		gauge     *prometheus.GaugeVec
		gaugeName = fmt.Sprintf("%v_%v", string(model.GetBaseOptions().Role), name)
		gaugeLbs  = append(labels, e.metricSpanID, e.metricStatus, e.metricStartTime, e.metricEndTime)
	)
	if obj, ok = e.metrics.Load(name); ok {
		if gauge, ok = obj.(*prometheus.GaugeVec); ok && gauge != nil {
			return gauge
		}
	}

	gauge = metrics.SetupGaugeVec(gaugeName, gaugeLbs...)
	e.metrics.Store(name, gauge)
	return gauge
}
