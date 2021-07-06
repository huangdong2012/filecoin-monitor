package metric

import (
	"contrib.go.opencensus.io/exporter/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
	"grandhelmsman/filecoin-monitor/metric/metrics"
	"net/http"
	"testing"
	"time"
)

const(
	rabbit = "amqp://root:root@localhost/"
)

func setupMetric() http.Handler {
	reg := prom.DefaultRegisterer.(*prom.Registry)
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "lotus",
		Registry:  reg,
		Gatherer:  InitGatherer(reg), //自定义Gather用于同时收集opencensus的metrics和我自己埋点的metrics
	})
	if err != nil {
		panic(err)
	}
	return exporter
}

//1.prometheus收集的方式
func TestMetric1(t *testing.T) {
	handler := setupMetric()
	Init(rabbit, &Options{
		Exchange:     "zdz.exchange.metric",
		RouteKey:     "*",
	})

	go func() {
		for range time.Tick(time.Second * 10) {
			metrics.Lotus.Test("label1", "label2", "label3").Inc()
			metrics.Lotus.Test2("label1", "label2", "label3").Set(float64(time.Now().Unix()))
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("/metrics", handler)
	if err := http.ListenAndServe(":3456", mux); err != nil {
		t.Fatal(err)
	}
}

//2.主动push到push gateway的方式
func TestMetric2(t *testing.T) {
	Init(rabbit, &Options{
		Exchange:     "zdz.exchange.metric",
		RouteKey:     "*",

		PushUrl:      "http://localhost:9091",
		PushJob:      "test-job",
		PushInterval: time.Second * 10,
	})

	metrics.Lotus.Test("label1", "label2", "label3").Inc()
	metrics.Lotus.Test2("label1", "label2", "label3").Set(float64(time.Now().Unix()))
	Push()

	time.Sleep(time.Second * 10)
}
