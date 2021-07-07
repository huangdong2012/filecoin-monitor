package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	prefixMiner = string(model.RoleMiner)
)

var (
	Miner = &minerMetrics{
		blockCount:     SetupCounterVec(naming(prefixMiner, "block_count")),
		nullRoundCount: SetupCounterVec(naming(prefixMiner, "null_round_count")),
	}
)

type minerMetrics struct {
	blockCount     *prometheus.CounterVec
	nullRoundCount *prometheus.CounterVec
}

func (m *minerMetrics) BlockCount() prometheus.Counter {
	return m.blockCount.WithLabelValues()
}

func (m *minerMetrics) NullRoundCount() prometheus.Counter {
	return m.nullRoundCount.WithLabelValues()
}
