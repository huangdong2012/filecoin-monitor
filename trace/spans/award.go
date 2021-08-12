package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewAwardSpan(ctx context.Context) (context.Context, *AwardSpan) {
	ct, span := setupSpan(ctx, "monitor-award")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &AwardSpan{&StatusSpan{span}}
}

type AwardSpan struct {
	*StatusSpan
}

func (s *AwardSpan) SetSystemActorAddr(addr string) {
	s.AddAttributes(trace.StringAttribute("system_actor_addr", addr))
}

func (s *AwardSpan) SetMinerID(mid string) {
	s.AddAttributes(trace.StringAttribute("miner_id", mid))
}

func (s *AwardSpan) SetEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("epoch", epoch))
}

func (s *AwardSpan) SetCID(cid string) {
	s.AddAttributes(trace.StringAttribute("cid", cid))
}

func (s *AwardSpan) SetMsgCount(count int) {
	s.AddAttributes(trace.Int64Attribute("msg_count", int64(count)))
}

func (s *AwardSpan) SetBaseFee(fee uint64) {
	s.AddAttributes(trace.Int64Attribute("base_fee", int64(fee)))
}

func (s *AwardSpan) SetPenalty(p uint64) {
	s.AddAttributes(trace.Int64Attribute("penalty", int64(p)))
}

func (s *AwardSpan) SetReward(r uint64) {
	s.AddAttributes(trace.Int64Attribute("reward", int64(r)))
}

func (s *AwardSpan) SetWinCount(count int64) {
	s.AddAttributes(trace.Int64Attribute("win_count", count))
}
