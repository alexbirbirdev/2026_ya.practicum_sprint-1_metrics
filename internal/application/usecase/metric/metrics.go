package metric

import (
	"errors"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	appErrors "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
)

type MetricUsecases struct {
	repo MetricsRepository
}

func NewMetricUsecases(
	repo MetricsRepository,
) *MetricUsecases {
	return &MetricUsecases{
		repo: repo,
	}
}

func (u *MetricUsecases) IncrementCounter(in dto.Metrics) (dto.Metrics, error) {
	metric, err := entity.NewCounterMetric(in.ID, *in.Delta)
	if err != nil {
		return dto.Metrics{}, err
	}
	existMetric, err := u.repo.FindMetricByNameAndType(metric.Name, vo.CounterMetricType)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return dto.Metrics{}, err
			}
			return dto.Metrics{
				ID:    metric.Name,
				MType: vo.CounterMetricType.String(),
				Delta: metric.CounterValue,
			}, nil
		}
		return dto.Metrics{}, err
	}
	newDelta, err := existMetric.IncrementCounterValue(metric.CounterValue)
	if err != nil {
		return dto.Metrics{}, err
	}

	err = u.repo.UpdateMetric(existMetric)
	if err != nil {
		return dto.Metrics{}, err
	}
	responseDTO := dto.Metrics{
		ID:    in.ID,
		MType: in.MType,
		Delta: newDelta,
	}

	return responseDTO, nil
}

func (u *MetricUsecases) UpdateGauge(in dto.Metrics) (dto.Metrics, error) {
	metric, err := entity.NewGaugeMetric(in.ID, *in.Value)
	if err != nil {
		return dto.Metrics{}, err
	}
	existMetric, err := u.repo.FindMetricByNameAndType(metric.Name, vo.MetricType(metric.Type))
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return dto.Metrics{}, err
			}
			return dto.Metrics{
				ID:    metric.Name,
				MType: vo.CounterMetricType.String(),
				Value: metric.GaugeValue,
			}, nil
		}
		return dto.Metrics{}, err
	}
	var newValue *float64
	if metric.GaugeValue != nil {
		newValue = metric.GaugeValue
	}
	val, err := existMetric.UpdateGaugeValue(newValue)
	if err != nil {
		return dto.Metrics{}, err
	}
	err = u.repo.UpdateMetric(existMetric)
	if err != nil {
		return dto.Metrics{}, err
	}
	responseDTO := dto.Metrics{
		ID:    in.ID,
		MType: in.MType,
		Value: val,
	}

	return responseDTO, nil
}

func (u *MetricUsecases) GetAllMetrics() ([]dto.Metrics, error) {
	var list []dto.Metrics

	metrics, err := u.repo.ListMetrics()
	if err != nil {
		return []dto.Metrics{}, err
	}
	for _, m := range metrics {
		metric := m.MapToDTO()
		list = append(list, metric)
	}
	return list, nil
}

func (u *MetricUsecases) GetMetricByNameAndType(metricName string, metricType vo.MetricType) (dto.Metrics, error) {
	existMetric, err := u.repo.FindMetricByNameAndType(metricName, vo.MetricType(metricType))
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			return dto.Metrics{}, appErrors.ErrNotFound
		}
		return dto.Metrics{}, err
	}

	metric := existMetric.MapToDTO()

	return metric, nil
}
