package metric

type MetricReader interface {
	Read() []Metric
}
