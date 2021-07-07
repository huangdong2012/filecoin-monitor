package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixLotus = string(model.RoleLotus)
)

var (
	Lotus = &lotusMetrics{
		test:  SetupCounterVec(naming(prefixLotus, "test")),
		test2: SetupGaugeVec(naming(prefixLotus, "test2")),
	}
)

type lotusMetrics struct {
	test  *prometheus.CounterVec
	test2 *prometheus.GaugeVec
}

func (m *lotusMetrics) Test() prometheus.Counter {
	return m.test.WithLabelValues()
}

func (m *lotusMetrics) Test2() prometheus.Gauge {
	return m.test2.WithLabelValues()
}
