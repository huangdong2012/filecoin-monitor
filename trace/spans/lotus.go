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

func (s *lotusSpan) Test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}

func (s *lotusSpan) SyncSubmitBlock(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/syncSubmitBlock")
}
