package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewWinningPostSpan(ctx context.Context) (context.Context, *WinningPostSpan) {
	ct, span := setupSpan(ctx, "monitor-winning-post")
	return ct, &WinningPostSpan{span}
}

type WinningPostSpan struct {
	*trace.Span
}

func (s *WinningPostSpan) Starting() {
	startingSpan(s.Span)
}

func (s *WinningPostSpan) Finish(err error) {
	finishSpan(s.Span, err)
}

func (s *WinningPostSpan) SetEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("epoch", epoch))
}

func (s *WinningPostSpan) SetBeacon(beacon string) {
	s.AddAttributes(trace.StringAttribute("beacon", beacon))
}

func (s *WinningPostSpan) SetProve(prove string) {
	s.AddAttributes(trace.StringAttribute("prove", prove))
}
