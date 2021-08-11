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

func (s *AwardSpan) SetMinerID(mid string) {
	s.AddAttributes(trace.StringAttribute("miner_id", mid))
}

func (s *AwardSpan) SetEpoch(epoch int64) {
	s.AddAttributes(trace.Int64Attribute("epoch", epoch))
}

func (s *AwardSpan) SetPenalty(p int64) {
	s.AddAttributes(trace.Int64Attribute("penalty", p))
}

func (s *AwardSpan) SetReward(r int64) {
	s.AddAttributes(trace.Int64Attribute("reward", r))
}

func (s *AwardSpan) SetWinCount(count int) {
	s.AddAttributes(trace.Int64Attribute("win_count", int64(count)))
}
