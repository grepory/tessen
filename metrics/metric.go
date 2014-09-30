package metrics

// A Metric is typically an association between a metric name and a measurement
// in time.
type Metric struct {
	Name      string
	Value     float64
	Timestamp int64
}
