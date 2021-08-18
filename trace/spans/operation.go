package spans

import (
	"context"
	"go.opencensus.io/trace"
)

func NewWorkerOperationSpan(ctx context.Context) (context.Context, *WorkerOperationSpan) {
	ct, span := setupSpan(ctx, "monitor-operation-worker")
	span.AddAttributes(trace.BoolAttribute(tagMetricEnable, true)) //导出metric
	return ct, &WorkerOperationSpan{&StatusSpan{span}}
}

type WorkerOperationSpan struct {
	*StatusSpan
}

func (s *WorkerOperationSpan) SetState(state bool) {
	s.AddAttributes(trace.BoolAttribute("state", state))
}

func (s *WorkerOperationSpan) SetWorkIP(ip string) {
	s.AddAttributes(trace.StringAttribute("work_ip", ip))
}

func (s *WorkerOperationSpan) SetWorkNo(no string) {
	s.AddAttributes(trace.StringAttribute("work_no", no))
}
