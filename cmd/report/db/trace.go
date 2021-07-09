package db

import (
	"grandhelmsman/filecoin-monitor/model"
	"grandhelmsman/filecoin-monitor/utils"
	"time"
)

func InsertSpan(span *model.Span) error {
	if span == nil {
		return nil
	}

	entity := &Span{
		ID:         span.ID,
		ParentID:   span.ParentID,
		TraceID:    span.TraceID,
		Service:    span.Service,
		Operation:  span.Operation,
		Tags:       utils.ToJsonWithoutError(span.Tags),
		Logs:       utils.ToJsonWithoutError(span.Logs),
		Duration:   span.Duration,
		Status:     span.Status,
		StartTime:  time.Unix(span.StartTime, 0),
		EndTime:    time.Unix(span.EndTime, 0),
		CreateTime: time.Now(),
	}

	return db.Create(entity).Error
}
