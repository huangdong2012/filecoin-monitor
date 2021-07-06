package trace

import (
	"errors"
	"fmt"
	"go.opencensus.io/trace"
)

var (
	options *Options
)

func Init(mq string, opt *Options) {
	if len(opt.Exchange) == 0 || len(opt.RouteKey) == 0 {
		panic("trace exchange or route-key invalid")
	}

	{
		options = opt
		initRabbit(mq)
	}

	trace.RegisterExporter(newExporter())
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
}

func parseSpan(sd *trace.SpanData) (*Span, error) {
	if sd == nil {
		return nil, errors.New("trace SpanData invalid")
	}

	span := &Span{
		ID:        sd.SpanID.String(),
		ParentID:  sd.ParentSpanID.String(),
		TraceID:   sd.TraceID.String(),
		Service:   "todo...",
		Operation: sd.Name,
		Tags:      make(map[string]string),
		Logs:      make(map[string]string),
		Duration:  sd.EndTime.Sub(sd.StartTime).Nanoseconds() / 1000,
		Status:    sd.Status.Code,
		StartTime: sd.StartTime.Unix(),
		EndTime:   sd.EndTime.Unix(),
	}
	for k, v := range sd.Attributes {
		span.Tags[k] = fmt.Sprintf("%v", v)
	}
	for _, ant := range sd.Annotations {
		span.Logs["message"] = ant.Message
		span.Logs["time"] = ant.Time.String()
		for k, v := range ant.Attributes {
			span.Logs[k] = fmt.Sprintf("%v", v)
		}
	}
	if sd.Status.Code != 0 {
		span.Logs["error"] = sd.Status.Message
	}
	return span, nil
}
