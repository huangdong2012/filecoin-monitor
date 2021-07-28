package trace

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"grandhelmsman/filecoin-monitor/utils"
)

var (
	options    *model.TraceOptions
	logger     *logrus.Entry
	spanLogger *utils.Logger
)

func Init(baseOpt *model.BaseOptions, traceOpt *model.TraceOptions) {
	model.InitBaseOptions(baseOpt)
	options = traceOpt

	log, err := utils.CreateLog(baseOpt.LogDir, baseOpt.LogTraceName, logrus.TraceLevel, true)
	if err != nil {
		panic(err)
	}
	logger = log.WithFields(logrus.Fields{
		"room-id":  baseOpt.RoomID,
		"miner-id": baseOpt.MinerID,
	})

	if len(options.SpanLogName) == 0 {
		options.SpanLogDir = "monitor-spans"
	}
	spanLogger, err = utils.CreateLog(options.SpanLogDir, options.SpanLogName, logrus.TraceLevel, true)

	trace.RegisterExporter(exp)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	//starting-handler
	spans.StartingHandler = exp.ExportSpan
}

func parseSpan(sd *trace.SpanData) (*model.Span, error) {
	if sd == nil {
		return nil, errors.New("trace SpanData invalid")
	}

	span := &model.Span{
		ID:        sd.SpanID.String(),
		ParentID:  sd.ParentSpanID.String(),
		TraceID:   sd.TraceID.String(),
		Service:   string(model.GetBaseOptions().PackageKind),
		Operation: sd.Name,
		Tags:      make(map[string]string),
		Logs:      make(map[string]string),
		Duration:  sd.EndTime.Sub(sd.StartTime).Seconds(),
		Status:    sd.Status.Code,
		StartTime: sd.StartTime.Unix(),
		EndTime:   sd.EndTime.Unix(),
	}
	for k, v := range sd.Attributes {
		span.Tags[k] = fmt.Sprintf("%v", v)
	}
	for _, ant := range sd.Annotations {
		span.Logs["message"] = ant.Message
		span.Logs["time"] = ant.Time.String()
		for k, v := range ant.Attributes {
			span.Logs[k] = fmt.Sprintf("%v", v)
		}
	}
	if sd.Status.Code != 0 {
		span.Tags["message"] = sd.Status.Message
		span.Logs["error"] = sd.Status.Message
	}
	return span, nil
}
