package main

type Metric struct {
	Name string // metric name in dotted form
	Timestamp uint64 // timestamp of collection in ms
	Samples []Sample // slice of samples relative to this Metrics' timestamp
}

type SampleType int 

const (
	GUAGE = iota
	COUNTER 
	DERIVE
	ABSOLUTE
)

type Sample struct {
	Name string // sample name, no dots
	Type SampleType // see above
	Value uint64
}

type MetricChannel chan Metric
