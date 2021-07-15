package spans

import (
	"context"
	"go.opencensus.io/trace"
	"time"
)

func NewWindowPostSpan(ctx context.Context) (context.Context, *WindowPostSpan) {
	ct, span := setupSpan(ctx, "monitor-window-post")
	return ct, &WindowPostSpan{span}
}

type WindowPostSpan struct {
	*trace.Span
}

func (s *WindowPostSpan) Starting(msg string) {
	startingSpan(s.Span, msg)
}

func (s *WindowPostSpan) Finish(err error) {
	finishSpan(s.Span, err)
}

func (s *SectorSpan) SetDeadline(deadline int) {
	s.AddAttributes(trace.Int64Attribute("deadline", int64(deadline)))
}

func (s *SectorSpan) SetPartitions(partitions string) {
	s.AddAttributes(trace.StringAttribute("partitions", partitions))
}

func (s *SectorSpan) SetPartitionCount(count int) {
	s.AddAttributes(trace.Int64Attribute("partition_count", int64(count)))
}

func (s *SectorSpan) SetSectorCount(count int) {
	s.AddAttributes(trace.Int64Attribute("sector_count", int64(count)))
}

func (s *SectorSpan) SetOpenTime(ot time.Time) {
	s.AddAttributes(trace.Int64Attribute("open_time", ot.Unix()))
}

func (s *SectorSpan) SetCloseTime(ct time.Time) {
	s.AddAttributes(trace.Int64Attribute("close_time", ct.Unix()))
}
