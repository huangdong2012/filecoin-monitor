package trace

import (
	"context"
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"

	"os"
	"testing"
)

const (
	service  = "test-srv"
	setupKey = "zdz"
)

func setupTrace() {
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")
	if len(agentEndpointURI) == 0 {
		panic("jaeger agent endpoint url invalid")
	}

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   service,
	})
	if err != nil {
		panic(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
}

func TestTrace(t *testing.T) {
	setupTrace()
	Init("amqp://root:root@localhost/", &Options{
		Exchange: "zdz.exchange.trace",
		RouteKey: "*",
		Service:  service,
	})

	ctx, span := trace.StartSpan(context.Background(), "/root")
	span.AddAttributes(trace.BoolAttribute(setupKey, true))
	span.AddAttributes(trace.StringAttribute("name", "aaa"))

	_, span1 := trace.StartSpan(ctx, "/sub1")
	span1.AddAttributes(trace.BoolAttribute(setupKey, true))
	span1.AddAttributes(trace.StringAttribute("name", "bbb"))
	span1.Annotate([]trace.Attribute{trace.StringAttribute("lang", "rust")}, "this is span1")
	span1.End()

	_, span2 := trace.StartSpan(ctx, "/sub2")
	span2.AddAttributes(trace.BoolAttribute(setupKey, true))
	span2.AddAttributes(trace.StringAttribute("name", "ccc"))
	span2.Annotate([]trace.Attribute{trace.StringAttribute("lang", "go")}, "this is span2")
	span2.End()

	span.End()
}
