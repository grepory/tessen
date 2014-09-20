package tessen

import "fmt"

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
}

// Collect is called at run-time by the class implementing a Collector
// plugin.
func (c *Collector) Collect() {
	for _, metric := range c.CollectFunc() {
		fmt.Printf("%s\t%f\t%d\n", metric.Name, metric.Value, metric.Timestamp)
	}
}
