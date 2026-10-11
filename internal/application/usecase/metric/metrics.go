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

func (u *MetricUsecases) IncrementCounter(in dto.MetricDTO) error {
	metric, err := entity.NewCounterMetric(in.Name, *in.CounterValue)
	if err != nil {
		return err
	}
	existMetric, err := u.repo.FindMetricByNameAndType(metric.Name, vo.CounterMetricType)
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err = existMetric.IncrementCounterValue(metric.CounterValue)
	if err != nil {
		return err
	}

	err = u.repo.UpdateMetric(existMetric)
	if err != nil {
		return err
	}

	return nil
}

func (u *MetricUsecases) UpdateGauge(in dto.MetricDTO) error {
	metric, err := entity.NewGaugeMetric(in.Name, *in.GaugeValue)
	if err != nil {
		return err
	}
	existMetric, err := u.repo.FindMetricByNameAndType(metric.Name, vo.MetricType(metric.Type))
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	var newValue *float64
	if metric.GaugeValue != nil {
		newValue = metric.GaugeValue
	}
	err = existMetric.UpdateGaugeValue(newValue)
	if err != nil {
		return err
	}
	err = u.repo.UpdateMetric(existMetric)
	if err != nil {
		return err
	}
	return nil
}

func (u *MetricUsecases) GetAllMetrics() ([]dto.MetricDTO, error) {
	var list []dto.MetricDTO

	metrics, err := u.repo.ListMetrics()
	if err != nil {
		return []dto.MetricDTO{}, err
	}
	for _, m := range metrics {
		metric := m.MapToDTO()
		list = append(list, metric)
	}
	return list, nil
}

func (u *MetricUsecases) GetMetricByNameAndType(metricName string, metricType vo.MetricType) (dto.MetricDTO, error) {
	existMetric, err := u.repo.FindMetricByNameAndType(metricName, vo.MetricType(metricType))
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			return dto.MetricDTO{}, appErrors.ErrNotFound
		}
		return dto.MetricDTO{}, err
	}

	metric := existMetric.MapToDTO()

	return metric, nil
}
