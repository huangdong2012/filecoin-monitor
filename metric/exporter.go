package metric

import (
	"grandhelmsman/filecoin-monitor/metric/metrics"
	"grandhelmsman/filecoin-monitor/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/push"
	"time"

	dto "github.com/prometheus/client_model/go"
)

var (
	exp = &exporter{
		exitC: make(chan bool),
		pushC: make(chan bool),
	}
)

type exporter struct {
	exitC chan bool
	pushC chan bool
}

func (e *exporter) start() {
	for {
		select {
		case <-e.exitC:
			utils.Info("metrics exporter exit loop")
			return
		case <-e.pushC:
		case <-time.After(options.PushInterval):
		}

		//send to push-gateway
		if err := push.New(options.PushUrl, options.PushJob).Gatherer(metrics.Registry).Push(); err != nil {
			utils.Error(fmt.Errorf("metrics exporter push error:%v", err.Error()))
		}

		if ms, err := metrics.InnerMetrics(); err != nil {
			utils.Error(fmt.Errorf("metrics exporter gather inner metrics error: %v", err.Error()))
		} else {
			e.export(ms)
		}
	}
}

func (e *exporter) stop() {
	select {
	case <-e.exitC:
	default:
		close(e.exitC)
	}
}

func (e *exporter) push() {
	select {
	case e.pushC <- true:
	case <-time.After(time.Millisecond * 200):
	}
}

//send to mq
func (e *exporter) export(ms []*dto.MetricFamily) {
	var (
		err  error
		data string
		out  = make([]*Metric, 0, 0)
	)
	for _, mf := range ms {
		if items := parseMetrics(mf); len(items) > 0 {
			out = append(out, items...)
		}
	}
	if data, err = utils.ToJson(out); err != nil {
		utils.Error(fmt.Errorf("marshal metrics to json error: %v", err))
		return
	}

	if err = sendToRabbit([]byte(data)); err != nil {
		utils.Error(fmt.Errorf("metric exporter send to mq error: %v", err.Error()))
	}
}
