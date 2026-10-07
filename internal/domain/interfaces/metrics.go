package interfaces

import "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"

type MetricsRepository interface {
	FindMetric(entity.Metric) (entity.Metric, error)
	CreateMetric(entity.Metric) error
	UpdateMetric(entity.Metric) error
}
