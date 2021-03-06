package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewSectorSpan(ctx context.Context) (context.Context, *SectorSpan) {
	ct, span := setupSpan(ctx, "monitor-sector")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &SectorSpan{&StatusSpan{span}}
}

type SectorSpan struct {
	*StatusSpan
}

func (s *SectorSpan) SetNumber(num int64) {
	s.AddAttributes(trace.Int64Attribute("number", num))
}

// 2K/8M/512M/32G/64G
func (s *SectorSpan) SetSize(size string) {
	s.AddAttributes(trace.StringAttribute("size", size))
}

// Empty
// WaitDeals
// AddPiece
// Packing
// GetTicket
// PreCommit1
// PreCommit2
// PreCommitting
// PreCommitWait
// WaitSeed
// Committing
// SubmitCommit
// CommitWait
// FinalizeSector
// Proving
func (s *SectorSpan) SetStep(step string) {
	s.AddAttributes(trace.StringAttribute("step", step))
}

func (s *SectorSpan) SetPath(path string) {
	s.AddAttributes(trace.StringAttribute("path", path))
}

func (s *SectorSpan) SetWorkIP(ip string) {
	s.AddAttributes(trace.StringAttribute("work_ip", ip))
}

func (s *SectorSpan) SetWorkNo(no string) {
	s.AddAttributes(trace.StringAttribute("work_no", no))
}
