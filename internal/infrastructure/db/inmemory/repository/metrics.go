package repository

import (
	"sync"

	interfaces "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/metric"
)

var _ interfaces.MetricsRepository = (*MetricsRepository)(nil)

type MetricsRepository struct {
	mu sync.Mutex
	db *models.InMemoryMetrics
}

func NewMetricsRepository(
	db *models.InMemoryMetrics,
) *MetricsRepository {
	return &MetricsRepository{
		db: db,
	}
}

func (r *MetricsRepository) ReplaceAllMetrics(metrics []entity.Metric) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.db.Metrics = make([]models.Metric, 0, len(metrics))
	for i := range metrics {
		r.db.Metrics = append(r.db.Metrics, *models.MapToModel(&metrics[i]))
	}
}

func (r *MetricsRepository) CreateMetric(metric entity.Metric) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	newMetric := models.MapToModel(&metric)
	r.db.Metrics = append(r.db.Metrics, *newMetric)
	return nil
}
func (r *MetricsRepository) UpdateMetric(metric entity.Metric) error {
	r.mu.Lock()
	defer r.mu.Unlock()

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
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []entity.Metric

	for _, metricModel := range r.db.Metrics {
		list = append(list, *metricModel.MapToEntity())
	}

	return list, nil
}

func (r *MetricsRepository) FindMetricByNameAndType(
	metricName string,
	metricType vo.MetricType,
) (entity.Metric, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	response := entity.Metric{}
	for _, inMemoryMetric := range r.db.Metrics {
		if inMemoryMetric.ID == metricName && vo.MetricType(inMemoryMetric.MType) == metricType {
			response = *inMemoryMetric.MapToEntity()
			return response, nil
		}
	}
	return response, errors.ErrNotFound
}
