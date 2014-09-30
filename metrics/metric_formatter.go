package metrics

type MetricFormatter interface {
	Format(Metric) string
}
