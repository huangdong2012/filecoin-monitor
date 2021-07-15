package spans

import (
	"context"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
)

const (
	setupKey = "zdz"
)

var (
	StartingHandler func(sd *trace.SpanData)
)

type StatusSpan struct {
	*trace.Span
}

func (s *StatusSpan) Starting(msg string) {
	startingSpan(s.Span, msg)
}

func (s *StatusSpan) Finish(err error) {
	finishSpan(s.Span, err)
}

func setupSpan(ctx context.Context, name string) (context.Context, *trace.Span) {
	ct, span := trace.StartSpan(ctx, name)
	span.AddAttributes(trace.BoolAttribute(setupKey, true))
	span.AddAttributes(trace.Int64Attribute("room_id", model.GetBaseOptions().RoomID))
	span.AddAttributes(trace.StringAttribute("host_ip", utils.IpAddr()))
	span.AddAttributes(trace.StringAttribute("miner", model.GetBaseOptions().Node)) // 如：to1000
	return ct, span
}

func startingSpan(span *trace.Span, msg string) {
	span.AddAttributes(trace.Int64Attribute("status", int64(model.WorkerStatus_Running)))
	span.AddAttributes(trace.StringAttribute("message", msg))
	if StartingHandler != nil {
		if sd := makeSpanData(span); sd != nil {
			StartingHandler(sd)
		}
	}
}

func finishSpan(span *trace.Span, err error) {
	if err == nil {
		span.AddAttributes(trace.Int64Attribute("status", int64(model.WorkerStatus_Finish)))
	} else {
		span.AddAttributes(trace.Int64Attribute("status", int64(model.WorkerStatus_Error)))
		span.AddAttributes(trace.StringAttribute("message", err.Error()))
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInternal,
			Message: err.Error(),
		})
	}
	span.End()
}

func makeSpanData(s *trace.Span) *trace.SpanData {
	//todo...

	return nil
}

func Verify(sd *trace.SpanData) bool {
	for k, _ := range sd.Attributes {
		if k == setupKey {
			return true
		}
	}
	return false
}
