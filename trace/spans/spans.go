package spans

import (
	"context"
	"go.opencensus.io/trace"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/utils"
	"strings"
)

const (
	tagSetupKey     = "zdz"
	tagRoomID       = "room_id"
	tagHostNo       = "host_no"
	tagHostIP       = "host_ip"
	tagMinerID      = "miner_id"
	tagStatus       = "status"
	tagMessage      = "message"
	tagMetricEnable = "metric-enable"
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
	initSpan(span)
	return ct, span
}

func setupSubSpan(spCtx *trace.SpanContext, name string) (context.Context, *trace.Span) {
	ct, span := trace.StartSpanWithRemoteParent(context.Background(), name, *spCtx)
	initSpan(span)
	return ct, span
}

func initSpan(span *trace.Span) {
	span.AddAttributes(trace.BoolAttribute(tagSetupKey, true))
	span.AddAttributes(trace.Int64Attribute(tagRoomID, model.GetBaseOptions().RoomID))
	span.AddAttributes(trace.StringAttribute(tagHostNo, model.GetBaseOptions().HostNo))
	span.AddAttributes(trace.StringAttribute(tagHostIP, utils.IpAddr()))
	span.AddAttributes(trace.StringAttribute(tagMinerID, model.GetBaseOptions().MinerID)) // 如：to1000
}

func startingSpan(span *trace.Span, msg string) {
	span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Running)))
	span.AddAttributes(trace.StringAttribute(tagMessage, msg))
	if StartingHandler != nil {
		if sd := span.Internal().MakeSpanData(); sd != nil {
			StartingHandler(sd)
		}
	}
}

func finishSpan(span *trace.Span, err error) {
	if err == nil {
		span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Finish)))
		span.AddAttributes(trace.StringAttribute(tagMessage, ""))
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

func IsLocalSpan(sd *trace.SpanData) bool {
	_, ok := sd.Attributes[tagSetupKey]
	return ok
}

func MetricEnable(tags map[string]string) bool {
	str, ok := tags[tagMetricEnable]
	return ok && strings.ToLower(str) == "true"
}
