package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewMineSpan(ctx context.Context) (context.Context, *MineSpan) {
	ct, span := setupSpan(ctx, "monitor-wining")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &MineSpan{&StatusSpan{span}}
}

type MineSpan struct {
	*StatusSpan
}

func (s *MineSpan) SetRound(round int64) {
	s.AddAttributes(trace.Int64Attribute("round", round))
}

func (s *MineSpan) SetNullRound(round int64) {
	s.AddAttributes(trace.Int64Attribute("null_round", round))
}

func (s *MineSpan) SetEligible(eligible bool) {
	s.AddAttributes(trace.BoolAttribute("eligible", eligible))
}

func (s *MineSpan) SetLookbackEpoch(epoch string) {
	s.AddAttributes(trace.StringAttribute("lookback_epoch", epoch))
}

func (s *MineSpan) SetBaseEpoch(epoch string) {
	s.AddAttributes(trace.StringAttribute("base_epoch", epoch))
}

func (s *MineSpan) SetBaseDeltaSeconds(bdf float64) {
	s.AddAttributes(trace.Float64Attribute("base_delta_seconds", bdf))
}

func (s *MineSpan) SetBeacon(beacon string) {
	s.AddAttributes(trace.StringAttribute("beacon", beacon))
}

func (s *MineSpan) SetBeaconEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("beacon_epoch", epoch))
}

func (s *MineSpan) SetLate(late bool) {
	s.AddAttributes(trace.BoolAttribute("late", late))
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

func (s *MineSpan) SetBaseInfoNil(isNil bool) {
	s.AddAttributes(trace.BoolAttribute("baseinfo_nil", isNil))
}

func (s *MineSpan) SetBaseInfoDuration(duration int64) {
	s.AddAttributes(trace.Int64Attribute("baseinfo_duration", duration))
}
