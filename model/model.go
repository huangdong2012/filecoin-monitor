package model

type Span struct {
	ID        string            `json:"id"`
	ParentID  string            `json:"parent_id"`
	TraceID   string            `json:"trace_id"`
	Service   string            `json:"service"`
	Operation string            `json:"operation"`
	Tags      map[string]string `json:"tags"`
	Logs      map[string]string `json:"logs"`
	Status    int32             `json:"status"`
	Duration  float64           `json:"duration"`
	StartTime int64             `json:"start_time"`
	EndTime   int64             `json:"end_time"`
}

type Metric struct {
	Name   string            `json:"name"`
	Desc   string            `json:"desc"`
	Value  float64           `json:"value"`
	Time   int64             `json:"time"`
	Labels map[string]string `json:"labels"`
}
