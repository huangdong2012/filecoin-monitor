package metrics

import (
	"grandhelmsman/filecoin-monitor/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	lotusSys = "lotus"
)

var (
	Lotus = &lotusMetrics{
		test: promauto.With(Registry).NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: lotusSys,
			Name:      "test",
		}, []string{
			instance,
			"label1",
			"label2",
			"label3",
		}),

		test2: promauto.With(Registry).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: lotusSys,
			Name:      "test2",
		}, []string{
			instance,
			"label1",
			"label2",
			"label3",
		}),
	}
)

type lotusMetrics struct {
	test  *prometheus.CounterVec
	test2 *prometheus.GaugeVec
}

func (m *lotusMetrics) Test(label1, label2, label3 string) prometheus.Counter {
	return m.test.WithLabelValues(utils.IpAddr(), label1, label2, label3)
}

func (m *lotusMetrics) Test2(label1, label2, label3 string) prometheus.Gauge {
	return m.test2.WithLabelValues(utils.IpAddr(), label1, label2, label3)
}
