package spans

import (
	"context"
	"go.opencensus.io/trace"
)

var (
	Storage = &storageSpan{}
)

type storageSpan struct {
}

func (s *storageSpan) test(ctx context.Context) (context.Context, *trace.Span) {
	return setupSpan(ctx, "/test")
}
