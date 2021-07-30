package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/utils"
	"strconv"
)

const (
	namespace     = "zdz"
	labelRoomID   = "room_id"
	labelInstance = "instance"
	labelHostNo   = "host_no"
	labelMinerID  = "miner_id"
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
			labelRoomID:   strconv.FormatInt(model.GetBaseOptions().RoomID, 10),
			labelInstance: utils.IpAddr(),
			labelHostNo:   model.GetBaseOptions().HostNo,
			labelMinerID:  model.GetBaseOptions().MinerID,
		},
	}
}
