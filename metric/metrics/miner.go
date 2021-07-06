package metrics

import (
	"grandhelmsman/filecoin-monitor/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	minerSys = "miner"
)

var (
	Miner = &minerMetrics{
		test: promauto.With(Registry).NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: minerSys,
			Name:      "test",
		}, []string{
			instance,
			"label1",
			"label2",
			"label3",
		}),
	}
)

type minerMetrics struct {
	test *prometheus.CounterVec
}

func (m *minerMetrics) Test(label1, label2, label3 string) prometheus.Counter {
	return m.test.WithLabelValues(utils.IpAddr(), label1, label2, label3)
}
