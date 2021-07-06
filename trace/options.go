package trace

type Options struct {
	Exchange string
	RouteKey string

	Service string
}

type Span struct {
	ID        string
	ParentID  string
	TraceID   string
	Service   string
	Operation string
	Tags      map[string]string
	Logs      map[string]string
	Duration  int64
	Status    int32
	StartTime int64
	EndTime   int64
}
