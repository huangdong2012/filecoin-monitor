package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewMineSpan(ctx context.Context) (context.Context, *MineSpan) {
	ct, span := setupSpan(ctx, "monitor-mine")
	return ct, &MineSpan{span}
}

type MineSpan struct {
	*trace.Span
}

func (s *MineSpan) Starting() {
	startingSpan(s.Span)
}

func (s *MineSpan) Finish(err error) {
	finishSpan(s.Span, err)
}

func (s *MineSpan) SetEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("epoch", epoch))
}

func (s *MineSpan) SetBeacon(beacon string) {
	s.AddAttributes(trace.StringAttribute("beacon", beacon))
}

func (s *MineSpan) SetTotalPower(power string) {
	s.AddAttributes(trace.StringAttribute("total_power", power))
}

func (s *MineSpan) SetMinerPower(power string) {
	s.AddAttributes(trace.StringAttribute("miner_power", power))
}

func (s *MineSpan) SetWinCount(count int) {
	s.AddAttributes(trace.Int64Attribute("win_count", int64(count)))
}

func (s *MineSpan) SetBlockCount(count int) {
	s.AddAttributes(trace.Int64Attribute("block_count", int64(count)))
}

func (s *MineSpan) SetMessage(msg string) {
	s.AddAttributes(trace.StringAttribute("message", msg))
}
