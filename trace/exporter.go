package trace

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/trace"
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
	//todo...

	//var (
	//	ok   bool
	//	name string
	//)
	//if span.Tags == nil {
	//	return nil
	//}
	//if name, ok = span.Tags[e.metricFlag]; !ok {
	//	return nil
	//}

	//metric := e.getMetric(name)
	//metric.With().Set(float64(span.Duration))

	return nil
}

func (e *Exporter) getMetric(name string) *prometheus.GaugeVec {
	//todo...
	//obj, ok := e.metrics.Load(name)
	//if !ok{
	//
	//}
	//
	//m := promauto.NewGaugeVec(nil)
	//m.
	return nil
}
