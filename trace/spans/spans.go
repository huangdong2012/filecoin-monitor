package spans

import (
	"context"
	"fmt"
	"go.opencensus.io/trace"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/utils"
	"strings"
	"time"
)

const (
	tagSetupKey     = "zdz"
	tagRoomID       = "room_id"
	tagHostNo       = "host_no"
	tagHostIP       = "host_ip"
	tagMinerID      = "miner_id"
	tagStatus       = "status"
	tagProcess      = "process"
	tagMessage      = "message"
	tagMetricEnable = "metric-enable"
)

var (
	ExportHandler func(sd *trace.SpanData)
)

type StatusSpan struct {
	*trace.Span
}

func (s *StatusSpan) Starting(msg string) {
	startingSpan(s.Span, msg)
}

func (s *StatusSpan) Process(name, msg string) {
	processSpan(s.Span, name, msg)
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
	span.AddAttributes(trace.StringAttribute(tagProcess, ""))
	span.AddAttributes(trace.Int64Attribute(tagStatus, int64(model.TaskStatus_Running)))
	span.AddAttributes(trace.StringAttribute(tagMessage, msg))
	if ExportHandler != nil {
		if sd := span.Internal().MakeSpanData(); sd != nil {
			ExportHandler(sd)
		}
	}
}

func processSpan(span *trace.Span, name, msg string) {
	if len(name) == 0 {
		return
	}
	span.AddAttributes(trace.StringAttribute(tagProcess, name))
	span.AddAttributes(trace.StringAttribute(fmt.Sprintf("%v_time", name), fmt.Sprintf("%v", time.Now().Unix())))
	if len(msg) > 0 {
		span.AddAttributes(trace.StringAttribute(fmt.Sprintf("%v_msg", name), msg))
	}
	if ExportHandler != nil {
		if sd := span.Internal().MakeSpanData(); sd != nil {
			ExportHandler(sd)
		}
	}
}

func finishSpan(span *trace.Span, err error) {
	span.AddAttributes(trace.StringAttribute(tagProcess, ""))
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
