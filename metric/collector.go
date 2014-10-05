package metric

import "fmt"

// A Collector reads and collects metrics, formats them, and prints them to
// STDOUT.
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
