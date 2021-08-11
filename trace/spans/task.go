package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewTaskSpan(ctx context.Context) (context.Context, *TaskSpan) {
	ct, span := setupSpan(ctx, "monitor-task")
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
