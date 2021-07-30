package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"huangdong2012/filecoin-monitor/model"
	"time"

	dto "github.com/prometheus/client_model/go"
)

var (
	exp = &exporter{
		exitC:     make(chan bool),
		pushC:     make(chan bool),
		pushColsC: make(chan []prometheus.Collector),
	}
)

type exporter struct {
	exitC     chan bool
	pushC     chan bool
	pushColsC chan []prometheus.Collector
}

func (e *exporter) start() {
	for {
		var (
			err          error
			cols         []prometheus.Collector
			gatherToMQ   prometheus.Gatherer = wrapperGather.inner
			gatherToProm prometheus.Gatherer = wrapperGather
		)
		select {
		case <-e.pushC: //push all
		case <-time.After(options.PushInterval): //push all
		case cols = <-e.pushColsC: //push cols
		case <-e.exitC:
			logger.Info("exporter exit loop")
			return
		}

		if len(cols) > 0 {
			reg := prometheus.NewRegistry()
			for _, c := range cols {
				if err = reg.Register(c); err != nil {
					logger.WithField("err", err).Error("exporter register error")
				}
			}
			gatherToMQ = reg
			gatherToProm = reg
		}

		//send to mq
		if options.ExportMetric != nil {
			if err := e.export(gatherToMQ); err != nil {
				logger.WithField("err", err).Error("exporter gather inner metrics error")
			}
		}

		//send to push-gateway
		if len(options.PushUrl) > 0 {
			if err := push.New(options.PushUrl, options.PushJob).Gatherer(gatherToProm).Push(); err != nil {
				logger.WithField("err", err).Error("exporter push error")
			}
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

func (e *exporter) pushAll() {
	select {
	case e.pushC <- true:
	case <-time.After(time.Millisecond * 200):
	}
}

func (e *exporter) pushCollectors(cols ...prometheus.Collector) {
	if len(cols) > 0 {
		select {
		case e.pushColsC <- cols:
		case <-time.After(time.Millisecond * 200):
		}
	}
}

//send to mq
func (e *exporter) export(gather prometheus.Gatherer) error {
	var (
		err error
		ms  []*dto.MetricFamily
		out = make([]*model.Metric, 0, 0)
	)
	if ms, err = gather.Gather(); err != nil {
		return err
	}
	if len(ms) == 0 {
		return nil
	}
	for _, mf := range ms {
		if items := parseMetrics(mf); len(items) > 0 {
			out = append(out, items...)
		}
	}
	if len(out) == 0 {
		return nil
	}
	if options.ExportMetric != nil {
		options.ExportMetric(out)
	}
	return nil
}
