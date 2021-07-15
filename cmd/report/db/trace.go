package db

import (
	"fmt"
	"gorm.io/gorm/clause"
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
		UpdateTime: time.Now(),
	}
	if span.Tags["status"] == fmt.Sprintf("%v", int32(model.WorkerStatus_Running)) {
		entity.Duration = 0
		entity.EndTime = entity.StartTime
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"tags":        entity.Tags,
			"logs":        entity.Logs,
			"update_time": time.Now(),
		}),
	}).Create(entity).Error
}
