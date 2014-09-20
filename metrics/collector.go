package metrics

import "bytes"

// A Metric is typically an association between a metric name and a measurement
// in time.
type Metric struct {
	Name      string
	Value     float64
	Timestamp int64
}

// A CollectFunc is a function that
type CollectFunc func() []Metric

// A Collector will collect metrics.
type Collector struct {
	CollectFunc
	Formatter Formatter
}

// NewCollector simplifies the creation of collectors by defaulting to the
// Graphite formatter.
func NewCollector(cf CollectFunc) *Collector {
	return &Collector{
		CollectFunc: cf,
		Formatter:   new(GraphiteFormatter),
	}
}

// Collect is called at run-time by the class implementing a Collector
// plugin.
func (c *Collector) Collect() string {
	var buffer bytes.Buffer
	for _, metric := range c.CollectFunc() {
		buffer.WriteString(c.Formatter.Format(metric))
	}
	return buffer.String()
}
