package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
	"strconv"
)

const (
	namespace = "zdz"
	roomID    = "room_id"
	instance  = "instance"
	miner     = "miner"
)

var (
	registry *prometheus.Registry //for zdz metrics to mq
)

func Init(reg *prometheus.Registry) {
	registry = reg
	{
		Lotus.init()
		Miner.init()
		Storage.init()
		Worker.init()
	}
}

func naming(prefix, name string) string {
	return fmt.Sprintf("%v_%v", prefix, name)
}

func SetupCounterVec(name string, labels ...string) *prometheus.CounterVec {
	return promauto.With(registry).NewCounterVec(prometheus.CounterOpts(setupMetricOptions(name)), labels)
}

func SetupGaugeVec(name string, labels ...string) *prometheus.GaugeVec {
	return promauto.With(registry).NewGaugeVec(prometheus.GaugeOpts(setupMetricOptions(name)), labels)
}

func setupMetricOptions(name string) prometheus.Opts {
	return prometheus.Opts{
		Namespace: namespace,
		Name:      name,
		ConstLabels: map[string]string{
			roomID:   strconv.FormatInt(model.GetBaseOptions().RoomID, 10),
			instance: utils.IpAddr(),
			miner:    model.GetBaseOptions().Node,
		},
	}
}
