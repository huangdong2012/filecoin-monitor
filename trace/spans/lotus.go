package spans

import (
	"context"
	"go.opencensus.io/trace"
)

var (
	Lotus = &lotusSpan{}
)

type lotusSpan struct {
}

func (s *lotusSpan) test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}
