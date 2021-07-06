package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

const (
	namespace = "zdz"
	instance = "instance"
)

var (
	Registry = prometheus.NewRegistry() //for zdz metrics to mq
)

func InnerMetrics() ([]*dto.MetricFamily, error) {
	return Registry.Gather()
}
