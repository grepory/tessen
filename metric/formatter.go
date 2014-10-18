package metric

// A MetricFormatter is responsible for turning a Metric struct into a string
// acceptable by a metrics backend. Examples:
// Graphite, OpenTSDB, Riemann
type Formatter interface {
	Format(Metric) string
}

func DefaultFormatter() Formatter {
	return &GraphiteFormatter{}
}
