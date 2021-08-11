package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewSectorManualSpan(ctx context.Context) (context.Context, *SectorManualSpan) {
	ct, span := setupSpan(ctx, "monitor-sector-manual")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &SectorManualSpan{&StatusSpan{span}}
}

type SectorManualSpan struct {
	*StatusSpan
}

func (s *SectorManualSpan) SetID(id string) {
	s.AddAttributes(trace.StringAttribute("id", id))
}

func (s *SectorManualSpan) SetMinerID(mid string) {
	s.AddAttributes(trace.StringAttribute("miner_id", mid))
}

func (s *SectorManualSpan) SetStep(step string) {
	s.AddAttributes(trace.StringAttribute("step", step))
}