package model

type Metric struct {
	Name   string
	Desc   string
	Value  float64
	Time   int64
	Labels map[string]string
}
