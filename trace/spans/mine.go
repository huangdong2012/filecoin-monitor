package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewMineSpan(ctx context.Context) (context.Context, *MineSpan) {
	ct, span := setupSpan(ctx, "monitor-wining")
	return ct, &MineSpan{&StatusSpan{span}}
}

type MineSpan struct {
	*StatusSpan
}

func (s *MineSpan) SetEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("epoch", epoch))
}

func (s *MineSpan) SetEligible(eligible int64) {
	s.AddAttributes(trace.Int64Attribute("eligible", eligible))
}

func (s *MineSpan) SetBeacon(beacon string) {
	s.AddAttributes(trace.StringAttribute("beacon", beacon))
}

func (s *MineSpan) SetLazy(lazy int64) {
	s.AddAttributes(trace.Int64Attribute("lazy", lazy))
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

func (s *MineSpan) SetMsgCount(count int) {
	s.AddAttributes(trace.Int64Attribute("msg_count", int64(count)))
}

func (s *MineSpan) SetBlockCount(count int) {
	s.AddAttributes(trace.Int64Attribute("block_count", int64(count)))
}

func (s *MineSpan) SetBaseInfoDuration(duration int) {
	s.AddAttributes(trace.Int64Attribute("baseinfo_duration", int64(duration)))
}
