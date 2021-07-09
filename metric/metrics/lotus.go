package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixLotus = string(model.RoleLotus)
)

var (
	Lotus = &lotusMetrics{}
)

type lotusMetrics struct {
	test  *prometheus.CounterVec
	test2 *prometheus.GaugeVec
}

func (m *lotusMetrics) init() {
	m.test = SetupCounterVec(naming(prefixLotus, "test"), "label1")
	m.test2 = SetupGaugeVec(naming(prefixLotus, "test2"), "label1")
}

func (m *lotusMetrics) Test(label1 string) prometheus.Counter {
	return m.test.WithLabelValues(label1)
}

func (m *lotusMetrics) Test2(label1 string) prometheus.Gauge {
	return m.test2.WithLabelValues(label1)
}
