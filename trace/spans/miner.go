package spans

import (
	"context"
	"go.opencensus.io/trace"
)

var (
	Miner = &minerSpan{}
)

type minerSpan struct {
}

func (s *minerSpan) test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}
