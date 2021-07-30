package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"huangdong2012/filecoin-monitor/model"
)

const (
	prefixStorage = string(model.PackageKind_Storage)
)

var (
	Storage = &storageMetrics{}
)

type storageMetrics struct {
	test *prometheus.CounterVec
}

func (m *storageMetrics) init() {
	m.test = SetupCounterVec(naming(prefixStorage, "test"))
}

func (m *storageMetrics) Test() prometheus.Counter {
	return m.test.WithLabelValues()
}
