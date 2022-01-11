package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewSectorFixSpan(ctx context.Context) (context.Context, *SectorFixSpan) {
	ct, span := setupSpan(ctx, "monitor-sector-fix")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &SectorFixSpan{&StatusSpan{span}}
}

type SectorFixSpan struct {
	*StatusSpan
}

func (s *SectorFixSpan) SetID(id string) {
	s.AddAttributes(trace.StringAttribute("id", id))
}

func (s *SectorFixSpan) SetMinerID(mid string) {
	s.AddAttributes(trace.StringAttribute("miner_id", mid))
}

//fix: remove/update
func (s *SectorFixSpan) SetFix(fix string) {
	s.AddAttributes(trace.StringAttribute("fix", fix))
}

func (s *SectorFixSpan) SetStep(step string) {
	s.AddAttributes(trace.StringAttribute("step", step))
}