package spans

import (
	"context"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/model"
)

const (
	setupKey = "zdz"
)

var (
	base *model.BaseOptions
)

func Init(baseOpt *model.BaseOptions) {
	base = baseOpt
}

func setupSpan(ctx context.Context, name string) (context.Context, *trace.Span) {
	ct, span := trace.StartSpan(ctx, name)
	span.AddAttributes(trace.BoolAttribute(setupKey, true))
	span.AddAttributes(trace.StringAttribute("node", base.Node)) // 如：to1000
	return ct, span
}

func Verify(sd *trace.SpanData) bool {
	for k, _ := range sd.Attributes {
		if k == setupKey {
			return true
		}
	}
	return false
}
