package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewLotusSpan(ctx context.Context) (context.Context, *NodeSpan) {
	ct, span := setupSpan(ctx, "monitor-lotus")
	return ct, &NodeSpan{span}
}

func NewMinerSpan(ctx context.Context) (context.Context, *NodeSpan) {
	ct, span := setupSpan(ctx, "monitor-miner")
	return ct, &NodeSpan{span}
}

func NewWorkerSpan(ctx context.Context) (context.Context, *NodeSpan) {
	ct, span := setupSpan(ctx, "monitor-worker")
	return ct, &NodeSpan{span}
}

type NodeSpan struct {
	*trace.Span
}

func (s *NodeSpan) SetInfo(info string) {
	s.AddAttributes(trace.StringAttribute("info", info))
}
