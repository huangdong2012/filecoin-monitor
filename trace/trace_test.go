package trace

import (
	"context"
	"grandhelmsman/filecoin-monitor/metric"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"time"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"

	"os"
	"testing"
)

const (
	setupKey = "zdz"
)

var (
	opt = &model.BaseOptions{
		Role:  model.Role_Miner,
		Node:  "t01000",
		MQUrl: "amqp://root:root@localhost/",
	}
)

func setupTrace() {
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")
	if len(agentEndpointURI) == 0 {
		panic("jaeger agent endpoint url invalid")
	}

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   string(opt.Role),
	})
	if err != nil {
		panic(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	Init(opt, &model.TraceOptions{
		Exchange: "zdz.exchange.trace",
		RouteKey: "*",
	})
	metric.Init(opt, &model.MetricOptions{
		Exchange: "zdz.exchange.metric",
		RouteKey: "*",

		PushUrl:      "http://localhost:9091",
		PushJob:      "test-job",
		PushInterval: time.Second * 10,
	})
}

func TestTrace(t *testing.T) {
	setupTrace()

	for i := 0; i < 3; i++ {
		ctx, span := trace.StartSpan(context.Background(), "/root")
		span.AddAttributes(trace.BoolAttribute(setupKey, true))
		span.AddAttributes(trace.StringAttribute("name", "aaa"))
		span.AddAttributes(trace.StringAttribute("metric", "root"))
		time.Sleep(time.Second)

		_, span1 := trace.StartSpan(ctx, "/sub1")
		span1.AddAttributes(trace.BoolAttribute(setupKey, true))
		span1.AddAttributes(trace.StringAttribute("name", "bbb"))
		span1.AddAttributes(trace.StringAttribute("metric", "sub1"))
		span1.Annotate([]trace.Attribute{trace.StringAttribute("lang", "rust")}, "this is span1")
		time.Sleep(time.Second)
		span1.End()

		_, span2 := trace.StartSpan(ctx, "/sub2")
		span2.AddAttributes(trace.BoolAttribute(setupKey, true))
		span2.AddAttributes(trace.StringAttribute("name", "ccc"))
		span2.Annotate([]trace.Attribute{trace.StringAttribute("lang", "go")}, "this is span2")
		time.Sleep(time.Second)
		span2.End()

		span.End()
	}

	time.Sleep(time.Second * 60)
}

func TestMineTrace(t *testing.T) {
	setupTrace()

	_, span := spans.NewMineSpan(context.Background())
	span.SetEpoch(1000)
	span.Starting()

	time.Sleep(time.Second * 5)
	span.SetBeacon("this is beacon")
	span.SetTotalPower("total power")
	span.SetMinerPower("miner power")
	span.SetWinCount(2)
	span.SetBlockCount(1)
	span.Finish(nil)
}
