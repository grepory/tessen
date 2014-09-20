package metrics

// FormatterFunc takes a tessen.Metric and turns it into a string that is ready
// for a metrics backend to eat.

// A MetricFormatter is responsible for turning a Metric struct into a string
// acceptable by a metrics backend. Examples:
// Graphite, OpenTSDB, Riemann
type Formatter interface {
	Format(Metric) string
}
