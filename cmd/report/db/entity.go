package db

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Span struct {
	ID         string    `gorm:"column:id;type:VARCHAR;size:256;" json:"id"`
	ParentID   string    `gorm:"column:parent_id;type:VARCHAR;size:256;" json:"parent_id"`
	TraceID    string    `gorm:"column:trace_id;type:VARCHAR;size:256;" json:"trace_id"`
	Service    string    `gorm:"column:service;type:VARCHAR;size:128;" json:"service"`
	Operation  string    `gorm:"column:operation;type:VARCHAR;size:256;" json:"operation"`
	Tags       string    `gorm:"column:tags;type:JSONB;" json:"tags"`
	Logs       string    `gorm:"column:logs;type:JSONB;" json:"logs"`
	Duration   float64   `gorm:"column:duration;type:NUMERIC;" json:"duration"`
	Status     int32     `gorm:"column:status;type:INT4;default:0;" json:"status"`
	StartTime  time.Time `gorm:"column:start_time;type:TIMESTAMPTZ;" json:"start_time"`
	EndTime    time.Time `gorm:"column:end_time;type:TIMESTAMPTZ;" json:"end_time"`
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
}

func (u *Span) TableName() string {
	return "spans"
}

func (u *Span) BeforeSave(tx *gorm.DB) error {
	return nil
}

type Metric struct {
	ID          int64          `gorm:"primary_key;column:id;type:INT8;" json:"id"`
	Name        string         `gorm:"column:name;type:VARCHAR;size:128;" json:"name"`
	Labels      string         `gorm:"column:labels;type:JSONB;" json:"labels"`
	Value       float64        `gorm:"column:value;type:NUMERIC;" json:"value"`
	Description sql.NullString `gorm:"column:description;type:VARCHAR;size:256;" json:"description"`
	MetricTime  time.Time      `gorm:"column:metric_time;type:TIMESTAMPTZ;" json:"metric_time"`
	CreateTime  time.Time      `gorm:"column:create_time;type:TIMESTAMPTZ;" json:"create_time"`
}

func (u *Metric) TableName() string {
	return "metrics"
}

func (u *Metric) BeforeSave(tx *gorm.DB) error {
	return nil
}
