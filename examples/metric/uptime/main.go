package main

// Regardless of whether or not this is a metric or status check, the main
// function is very small.

import (
	"github.com/grepory/tessen/metric"
	"github.com/grepory/tessen/metric/formatter"
)

func main() {
	collector := metric.Collector{
		Reader:    &UptimeReader{},
		Formatter: formatter.DefaultFormatter(),
	}
	collector.Run()
}
