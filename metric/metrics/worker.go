package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixWorker = string(model.Role_Worker)
)

var (
	Worker = &workerMetrics{}
)

type workerMetrics struct {
	test *prometheus.CounterVec
}

func (m *workerMetrics) init() {
	m.test = SetupCounterVec(naming(prefixWorker, "test"))
}

func (m *workerMetrics) Test() prometheus.Counter {
	return m.test.WithLabelValues()
}
