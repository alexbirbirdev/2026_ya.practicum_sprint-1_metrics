package interfaces

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type MetricsRepository interface {
	CreateMetric(entity.Metric) error
	UpdateMetric(entity.Metric) error
	ListMetrics() ([]entity.Metric, error)
	FindMetricByNameAndType(metricName string, metricType vo.MetricType) (entity.Metric, error)
}
