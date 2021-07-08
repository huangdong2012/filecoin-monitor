package metric

import "github.com/prometheus/client_golang/prometheus"

type Scope struct {
	collectors []prometheus.Collector
}

func NewScope() *Scope {
	return &Scope{
		collectors: make([]prometheus.Collector, 0, 0),
	}
}

func (s *Scope) Add(cols ...prometheus.Collector) *Scope {
	s.collectors = append(s.collectors, cols...)
	return s
}

func (s *Scope) Push() {
	exp.pushCollectors(s.collectors...)
}
