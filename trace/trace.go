package trace

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"huangdong2012/filecoin-monitor/model"
	"huangdong2012/filecoin-monitor/trace/spans"
	"huangdong2012/filecoin-monitor/utils"
)

var (
	options    *model.TraceOptions
	logger     *logrus.Entry
	spanLogger *utils.Logger
)

func Init(baseOpt *model.BaseOptions, traceOpt *model.TraceOptions) {
	model.InitBaseOptions(baseOpt)
	options = traceOpt

	log, err := utils.CreateLog(baseOpt.LogDir, baseOpt.LogTraceName, logrus.TraceLevel, baseOpt.LogToStdout)
	if err != nil {
		panic(err)
	}
	fields := logrus.Fields{"room-id": baseOpt.RoomID}
	if len(baseOpt.MinerID) > 0 {
		fields["miner-id"] = baseOpt.MinerID
	}
	logger = log.WithFields(fields)

	if len(options.SpanLogName) == 0 {
		options.SpanLogName = fmt.Sprintf("%v.monitor-span", string(baseOpt.PackageKind))
	}
	spanLogger, err = utils.CreateLog(options.SpanLogDir, options.SpanLogName, logrus.TraceLevel, baseOpt.LogToStdout)

	trace.RegisterExporter(exp)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	//starting-handler
	spans.ExportHandler = exp.ExportSpan
	logger.Info("monitor-trace starting...")
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
