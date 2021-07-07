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

func (s *minerSpan) Test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}

func (s *minerSpan) MineLoop(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/mineLoop")
}

func (s *minerSpan) MineOne(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/mineOne")
}
