package spans

import (
	"context"
	"go.opencensus.io/trace"
)

var (
	Worker = &workerSpan{}
)

type workerSpan struct {
}

func (s *workerSpan) Test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}
