package usecase

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/interfaces"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
)

type MetricUsecases struct {
	repo interfaces.MetricsRepository
}

func NewMetricUsecases(
	repo interfaces.MetricsRepository,
) *MetricUsecases {
	return &MetricUsecases{
		repo: repo,
	}
}

func (u *MetricUsecases) IncrementCounter(in dto.UpdateMetricsRequest) error {
	metric, err := entity.NewCounterMetric(in.Name, *in.CounterValue)
	if err != nil {
		return err
	}
	existMetric, err := u.repo.FindMetric(*metric)
	if err != nil {
		if err == errors.NotFound {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err = existMetric.IncrementCounterValue(*metric.CounterValue)
	if err != nil {
		return err
	}
	err = u.repo.UpdateMetric(existMetric)
	if err != nil {
		return err
	}
	return nil
}

func (u *MetricUsecases) UpdateGauge(in dto.UpdateMetricsRequest) error {
	metric, err := entity.NewGaugeMetric(in.Name, *in.GaugeValue)
	if err != nil {
		return err
	}
	existMetric, err := u.repo.FindMetric(*metric)
	if err != nil {
		if err == errors.NotFound {
			err = u.repo.CreateMetric(*metric)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	var newValue float64
	if metric.GaugeValue != nil {
		newValue = *metric.GaugeValue
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
