package formatter

import (
	"fmt"

	"github.com/grepory/tessen/metrics"
)

type GraphiteFormatter struct {
}

// Format returns a Graphite metric string given a Tessen.Metric object.
func (f *GraphiteFormatter) Format(metric metrics.Metric) string {
	return fmt.Sprintf("%s\t%f\t%d\n", metric.Name, metric.Value, metric.Timestamp)
}