package metrics

type MetricReader interface {
	Read() []Metric
}
