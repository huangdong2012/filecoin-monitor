package metrics

import (
	"grandhelmsman/filecoin-monitor/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	storageSys = "storage"
)

var (
	Storage = &storageMetrics{
		test: promauto.With(Registry).NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: storageSys,
			Name:      "test",
		}, []string{
			instance,
			"label1",
			"label2",
			"label3",
		}),
	}
)

type storageMetrics struct {
	test *prometheus.CounterVec
}

func (m *storageMetrics) Test(label1, label2, label3 string) prometheus.Counter {
	return m.test.WithLabelValues(utils.IpAddr(), label1, label2, label3)
}
