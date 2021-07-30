package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"huangdong2012/filecoin-monitor/model"
)

const (
	prefixWorker = string(model.PackageKind_Worker)
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
