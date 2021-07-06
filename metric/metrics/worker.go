package metrics

import (
	"grandhelmsman/filecoin-monitor/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	workerSys = "worker"
)

var (
	Worker = &workerMetrics{
		test: promauto.With(Registry).NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: workerSys,
			Name:      "test",
		}, []string{
			instance,
			"label1",
			"label2",
			"label3",
		}),
	}
)

type workerMetrics struct {
	test *prometheus.CounterVec
}

func (m *workerMetrics) Test(label1, label2, label3 string) prometheus.Counter {
	return m.test.WithLabelValues(utils.IpAddr(), label1, label2, label3)
}
