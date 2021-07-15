package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewSectorSpan(ctx context.Context) (context.Context, *SectorSpan) {
	ct, span := setupSpan(ctx, "monitor-sector")
	return ct, &SectorSpan{span}
}

type SectorSpan struct {
	*trace.Span
}

func (s *SectorSpan) Starting(msg string) {
	startingSpan(s.Span, msg)
}

func (s *SectorSpan) Finish(err error) {
	finishSpan(s.Span, err)
}

func (s *SectorSpan) SetNumber(num int64) {
	s.AddAttributes(trace.Int64Attribute("number", num))
}

/*
	2K
	8M
	512M
	32G
	64G
*/
func (s *SectorSpan) SetSize(size string) {
	s.AddAttributes(trace.StringAttribute("size", size))
}

/*
	Empty
	WaitDeals
	AddPiece
	Packing
	GetTicket
	PreCommit1
	PreCommit2
	PreCommitting
	PreCommitWait
	WaitSeed
	Committing
	SubmitCommit
	CommitWait
	FinalizeSector
	Proving
*/
func (s *SectorSpan) SetStep(step string) {
	s.AddAttributes(trace.StringAttribute("step", step))
}
