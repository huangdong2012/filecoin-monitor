package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixStorage = string(model.RoleStorage)
)

var (
	Storage = &storageMetrics{
		test: SetupCounterVec(naming(prefixStorage, "test")),
	}
)

type storageMetrics struct {
	test *prometheus.CounterVec
}

func (m *storageMetrics) Test() prometheus.Counter {
	return m.test.WithLabelValues()
}
