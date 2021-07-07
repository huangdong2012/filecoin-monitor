package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixWorker = string(model.RoleWorker)
)

var (
	Worker = &workerMetrics{
		test: SetupCounterVec(naming(prefixWorker, "test")),
	}
)

type workerMetrics struct {
	test *prometheus.CounterVec
}

func (m *workerMetrics) Test() prometheus.Counter {
	return m.test.WithLabelValues()
}
