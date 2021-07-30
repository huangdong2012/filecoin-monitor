module huangdong2012/filecoin-monitor

go 1.16

replace go.opencensus.io => github.com/huangdong2012/opencensus-go v0.0.0-20210728074244-7f6b340fc394

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/client_model v0.2.0
	github.com/sirupsen/logrus v1.6.0
	go.opencensus.io v0.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
