package repository

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/interfaces"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/metric"
)

var _ interfaces.MetricsRepository = (*MetricsRepository)(nil)

type MetricsRepository struct {
	db *models.InMemoryMetrics
}

func NewMetricsRepository(
	db *models.InMemoryMetrics,
) *MetricsRepository {
	return &MetricsRepository{
		db: db,
	}
}
func (r *MetricsRepository) CreateMetric(metric entity.Metric) error {
	newMetric := models.MapToModel(&metric)
	r.db.Metrics = append(r.db.Metrics, *newMetric)
	return nil
}
func (r *MetricsRepository) UpdateMetric(metric entity.Metric) error {
	for i := range r.db.Metrics {
		inMemoryMetric := &r.db.Metrics[i]
		if inMemoryMetric.ID == metric.Name && inMemoryMetric.MType == metric.Type {
			r.db.Metrics[i] = *models.MapToModel(&metric)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (r *MetricsRepository) ListMetrics() ([]entity.Metric, error) {
	var list []entity.Metric

	for _, metricModel := range r.db.Metrics {
		list = append(list, *metricModel.MapToEntity())
	}

	return list, nil
}

func (r *MetricsRepository) FindMetricByNameAndType(metricName string, metricType vo.MetricType) (entity.Metric, error) {
	response := entity.Metric{}
	for _, inMemoryMetric := range r.db.Metrics {
		if inMemoryMetric.ID == metricName && vo.MetricType(inMemoryMetric.MType) == metricType {
			response = *inMemoryMetric.MapToEntity()
			return response, nil
		}
	}
	return response, errors.ErrNotFound
}
