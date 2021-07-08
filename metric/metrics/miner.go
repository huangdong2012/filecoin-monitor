package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixMiner = string(model.RoleMiner)
)

var (
	Miner = &minerMetrics{}
)

type minerMetrics struct {
	blockCount     *prometheus.CounterVec
	nullRoundCount *prometheus.CounterVec
}

func (m *minerMetrics) init() {
	m.blockCount = SetupCounterVec(naming(prefixMiner, "block_count"))
	m.nullRoundCount = SetupCounterVec(naming(prefixMiner, "null_round_count"))
}

func (m *minerMetrics) BlockCount() prometheus.Counter {
	return m.blockCount.WithLabelValues()
}

func (m *minerMetrics) NullRoundCount() prometheus.Counter {
	return m.nullRoundCount.WithLabelValues()
}
