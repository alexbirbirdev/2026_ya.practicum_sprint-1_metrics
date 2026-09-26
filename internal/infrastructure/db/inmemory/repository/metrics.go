package repository

import (
	"log/slog"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/interfaces"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models"
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

func (r *MetricsRepository) FindMetric(metric entity.Metric) (entity.Metric, error) {
	response := entity.Metric{}
	for _, inMemoryMetric := range r.db.Metrics {
		if inMemoryMetric.ID == metric.Name && inMemoryMetric.MType == metric.Type {
			response = *inMemoryMetric.MapToEntity()
			return response, nil
		}
	}
	db := r.db.Metrics
	slog.Info("DB:", db)
	return response, errors.ErrNotFound
}
func (r *MetricsRepository) CreateMetric(metric entity.Metric) error {
	newMetric := models.MapToModel(&metric)
	r.db.Metrics = append(r.db.Metrics, *newMetric)
	db := r.db.Metrics
	slog.Info("DB:", db)
	return nil
}
func (r *MetricsRepository) UpdateMetric(metric entity.Metric) error {
	for _, inMemoryMetric := range r.db.Metrics {
		if inMemoryMetric.ID == metric.Name {
			inMemoryMetricDomain := inMemoryMetric.MapToEntity()
			if inMemoryMetricDomain.Type == vo.Counter {
				inMemoryMetricDomain.IncrementCounterValue(*metric.CounterValue)
				return nil
			}
			inMemoryMetricDomain.UpdateGaugeValue(*metric.GaugeValue)
			inMemoryMetric = *models.MapToModel(inMemoryMetricDomain)
			return nil
		}
	}
	return nil
}
