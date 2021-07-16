package spans

import (
	"context"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
	"strings"
)

const (
	tagSetupKey     = "zdz"
	tagMetricEnable = "metric-enable"
	tagRoomID       = "room_id"
	tagHostIP       = "host_ip"
	tagMinerID      = "miner_id"
	tagStatus       = "status"
	tagMessage      = "message"
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
	span.AddAttributes(trace.BoolAttribute(tagSetupKey, true))
	span.AddAttributes(trace.Int64Attribute(tagRoomID, model.GetBaseOptions().RoomID))
	span.AddAttributes(trace.StringAttribute(tagHostIP, utils.IpAddr()))
	span.AddAttributes(trace.StringAttribute(tagMinerID, model.GetBaseOptions().MinerID)) // 如：to1000
	return ct, span
}

func startingSpan(span *trace.Span, msg string) {
	span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Running)))
	span.AddAttributes(trace.StringAttribute(tagMessage, msg))
	if StartingHandler != nil {
		if sd := makeSpanData(span); sd != nil {
			StartingHandler(sd)
		}
	}
}

func finishSpan(span *trace.Span, err error) {
	if err == nil {
		span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Finish)))
	} else {
		span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Error)))
		span.AddAttributes(trace.StringAttribute(tagMessage, err.Error()))
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
		if k == tagSetupKey {
			return true
		}
	}
	return false
}

func MetricEnable(tags map[string]string) bool {
	if str, ok := tags[tagMetricEnable]; ok && strings.ToLower(str) == "false" {
		return false
	}
	return true
}
