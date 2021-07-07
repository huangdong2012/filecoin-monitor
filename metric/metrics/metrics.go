package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	dto "github.com/prometheus/client_model/go"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
)

const (
	namespace = "zdz"
	instance  = "instance"
	node      = "node"
)

var (
	base     *model.BaseOptions
	registry = prometheus.NewRegistry() //for zdz metrics to mq
)

func Init(baseOpt *model.BaseOptions) {
	base = baseOpt
}

func Registry() *prometheus.Registry {
	return registry
}

func InnerMetrics() ([]*dto.MetricFamily, error) {
	return registry.Gather()
}

func naming(prefix, name string) string {
	return fmt.Sprintf("%v_%v", prefix, name)
}

func SetupCounterVec(name string, labels ...string) *prometheus.CounterVec {
	return promauto.With(registry).NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		ConstLabels: map[string]string{
			instance: utils.IpAddr(),
			node:     base.Node,
		},
	}, labels)
}

func SetupGaugeVec(name string, labels ...string) *prometheus.GaugeVec {
	return promauto.With(registry).NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		ConstLabels: map[string]string{
			instance: utils.IpAddr(),
			node:     base.Node,
		},
	}, labels)
}
