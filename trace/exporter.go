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
)

func newExporter() *Exporter {
	return &Exporter{
		metricFlag: "metric",
		metrics:    &sync.Map{},
	}
}

type Exporter struct {
	metricFlag string
	metrics    *sync.Map
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

	gauge = e.getMetric(name, span.Tags)
	gauge.With(span.Tags).Set(span.Duration)
	metric.Push()

	return nil
}

func (e *Exporter) getMetric(name string, tags map[string]string) *prometheus.GaugeVec {
	var (
		ok    bool
		obj   interface{}
		gauge *prometheus.GaugeVec
	)
	if obj, ok = e.metrics.Load(name); ok {
		if gauge, ok = obj.(*prometheus.GaugeVec); ok && gauge != nil {
			return gauge
		}
	}

	gauge = metrics.SetupGaugeVec(fmt.Sprintf("%v_%v", string(model.GetBaseOptions().Role), name), utils.GetKeys(tags)...)
	e.metrics.Store(name, gauge)
	return gauge
}
