package model

type Span struct {
	ID        string
	ParentID  string
	TraceID   string
	Service   string
	Operation string
	Tags      map[string]string
	Logs      map[string]string
	Status    int32
	Duration  float64
	StartTime int64
	EndTime   int64
}

