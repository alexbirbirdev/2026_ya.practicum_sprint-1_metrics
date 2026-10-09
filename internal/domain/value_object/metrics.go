package valueobject

type MetricType string

var (
	CounterMetricType MetricType = "counter"
	GaugeMetricType   MetricType = "gauge"
)

func (m MetricType) String() string {
	return string(m)
}

func (m MetricType) Validate() bool {
	if m != "" {
		switch m {
		case CounterMetricType, GaugeMetricType:
			return true
		default:
			return false
		}
	}
	return false
}
