package trace

import (
	"fmt"
	"go.opencensus.io/trace"
	"grandhelmsman/filecoin-monitor/trace/spans"
	"grandhelmsman/filecoin-monitor/utils"
)

func newExporter() *Exporter {
	return &Exporter{}
}

type Exporter struct {
}

//send to mq
func (e *Exporter) ExportSpan(sd *trace.SpanData) {
	var (
		err  error
		span *Span
		data string
	)
	if !spans.Verify(sd) {
		return
	}
	if span, err = parseSpan(sd); err != nil {
		utils.Error(fmt.Errorf("parse span data error: %v", err))
		return
	}
	if data, err = utils.ToJson(span); err != nil {
		utils.Error(fmt.Errorf("marshal span to json error: %v", err))
		return
	}

	if err = sendToRabbit([]byte(data)); err != nil {
		utils.Error(fmt.Errorf("trace exporter send to mq error: %v", err.Error()))
	}
}
