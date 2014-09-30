package main

import (
	"github.com/grepory/tessen/formatter"
	"github.com/grepory/tessen/metrics"
)

func main() {
	collector := metrics.Collector{
		Reader:    &UptimeReader{},
		Formatter: formatter.DefaultFormatter(),
	}
	collector.Run()
}
