package metrics

import "fmt"

// A Collector will collect metrics.
type Collector struct {
	Reader    MetricReader
	Formatter MetricFormatter
}

func (c *Collector) Run() {
	metricsCollected := c.Reader.Read()
	for _, metric := range metricsCollected {
		fmt.Println(c.Formatter.Format(metric))
	}
}
