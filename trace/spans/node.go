package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewLotusSpan(ctx context.Context) (context.Context, *NodeSpan) {
	return newNodeSpan(ctx, "monitor-lotus")
}

func NewMinerSpan(ctx context.Context) (context.Context, *NodeSpan) {
	return newNodeSpan(ctx, "monitor-miner")
}

func NewWorkerSpan(ctx context.Context) (context.Context, *NodeSpan) {
	return newNodeSpan(ctx, "monitor-worker")
}

func NewStorageSpan(ctx context.Context) (context.Context, *NodeSpan) {
	return newNodeSpan(ctx, "monitor-storage")
}

func newNodeSpan(ctx context.Context, name string) (context.Context, *NodeSpan) {
	ct, span := setupSpan(ctx, name)
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, false)) //不导出metric
	return ct, &NodeSpan{span}
}

type NodeSpan struct {
	*trace.Span
}

func (s *NodeSpan) SetInfo(info string) {
	s.AddAttributes(trace.StringAttribute("info", info))
}
