package interfaces

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type MockMetricsRepository struct {
	CreateMetricFn            func(entity.Metric) error
	UpdateMetricFn            func(entity.Metric) error
	ListMetricsFn             func() ([]entity.Metric, error)
	FindMetricByNameAndTypeFn func(metricName string, metricType vo.MetricType) (entity.Metric, error)
}

func (m *MockMetricsRepository) CreateMetric(metric entity.Metric) error {
	if m.CreateMetricFn == nil {
		return nil
	}
	return m.CreateMetricFn(metric)
}
func (m *MockMetricsRepository) UpdateMetric(metric entity.Metric) error {
	if m.UpdateMetricFn == nil {
		return nil
	}
	return m.UpdateMetricFn(metric)
}
func (m *MockMetricsRepository) ListMetrics() ([]entity.Metric, error) {
	if m.ListMetricsFn == nil {
		return nil, nil
	}
	return m.ListMetricsFn()
}
func (m *MockMetricsRepository) FindMetricByNameAndType(
	metricName string,
	metricType vo.MetricType,
) (entity.Metric, error) {
	if m.FindMetricByNameAndTypeFn == nil {
		return entity.Metric{}, nil
	}
	return m.FindMetricByNameAndTypeFn(metricName, metricType)
}
