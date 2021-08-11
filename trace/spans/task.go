package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewTaskSpan(spCtx *trace.SpanContext) (context.Context, *TaskSpan) {
	name := "monitor-task"
	ct, span := setupSpan(context.Background(), name)
	if spCtx != nil {
		ct, span = setupSubSpan(spCtx, name)
	}
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &TaskSpan{&StatusSpan{span}}
}

type TaskSpan struct {
	*StatusSpan
}

func (s *TaskSpan) SetMinerID(mid string) {
	s.AddAttributes(trace.StringAttribute("miner_id", mid))
}

func (s *TaskSpan) SetType(typ string) {
	s.AddAttributes(trace.StringAttribute("type", typ))
}

func (s *TaskSpan) SetWorkIP(ip string) {
	s.AddAttributes(trace.StringAttribute("work_ip", ip))
}

func (s *TaskSpan) SetWorkNo(no string) {
	s.AddAttributes(trace.StringAttribute("work_no", no))
}

func (s *TaskSpan) SetMaxTaskCount(count int64) {
	s.AddAttributes(trace.Int64Attribute("max_task_count", count))
}

func (s *TaskSpan) SetWindowPostEnable(enable bool) {
	s.AddAttributes(trace.BoolAttribute("window_post_enable", enable))
}

func (s *TaskSpan) SetWinningPostEnable(enable bool) {
	s.AddAttributes(trace.BoolAttribute("winning_post_enable", enable))
}
