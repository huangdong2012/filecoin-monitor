package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewMPoolSpan(ctx context.Context) (context.Context, *MPoolSpan) {
	ct, span := setupSpan(ctx, "monitor-mpool")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, false)) //不导出metric
	return ct, &MPoolSpan{span}
}

type MPoolSpan struct {
	*trace.Span
}

func (s *MPoolSpan) SetInfo(info string) {
	s.AddAttributes(trace.StringAttribute("info", info))
}
