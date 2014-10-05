package metric

type MetricFormatter interface {
	Format(Metric) string
}
