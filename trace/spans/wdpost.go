package spans

import (
	"context"
	"go.opencensus.io/trace"
	"time"
)

func NewWindowPostSpan(ctx context.Context) (context.Context, *WindowPostSpan) {
	ct, span := setupSpan(ctx, "monitor-wdpost")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &WindowPostSpan{&StatusSpan{span}}
}

type WindowPostSpan struct {
	*StatusSpan
}

func (s *WindowPostSpan) SetDeadline(deadline int) {
	s.AddAttributes(trace.Int64Attribute("deadline", int64(deadline)))
}

func (s *WindowPostSpan) SetPartitions(partitions string) {
	s.AddAttributes(trace.StringAttribute("partitions", partitions))
}

func (s *WindowPostSpan) SetPartitionCount(count int) {
	s.AddAttributes(trace.Int64Attribute("partition_count", int64(count)))
}

func (s *WindowPostSpan) SetSectorCount(count int) {
	s.AddAttributes(trace.Int64Attribute("sector_count", int64(count)))
}

func (s *WindowPostSpan) SetSkipCount(count int) {
	s.AddAttributes(trace.Int64Attribute("skip_count", int64(count)))
}

func (s *WindowPostSpan) SetOpenTime(ot time.Time) {
	s.AddAttributes(trace.Int64Attribute("open_time", ot.Unix()))
}

func (s *WindowPostSpan) SetCloseTime(ct time.Time) {
	s.AddAttributes(trace.Int64Attribute("close_time", ct.Unix()))
}

func (s *WindowPostSpan) SetHeight(h int64) {
	s.AddAttributes(trace.Int64Attribute("height", h))
}

func (s *WindowPostSpan) SetRand(r string) {
	s.AddAttributes(trace.StringAttribute("rand", r))
}
