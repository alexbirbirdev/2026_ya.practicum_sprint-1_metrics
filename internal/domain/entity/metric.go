package entity

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type Metric struct {
	Name         string
	Type         string
	CounterValue *int64
	GaugeValue   *float64
}

func (m *Metric) IncrementCounterValue(value *int64) error {
	if m.CounterValue != nil {
		*m.CounterValue += *value
	}
	return nil
}

func (m *Metric) UpdateGaugeValue(value *float64) error {
	if m.GaugeValue != nil {
		*m.GaugeValue = *value
	}
	return nil
}

func NewCounterMetric(
	name string,
	value int64,
) (*Metric, error) {
	return &Metric{
		Name:         name,
		Type:         vo.CounterMetricType.String(),
		CounterValue: &value,
	}, nil
}
func NewGaugeMetric(
	name string,
	value float64,
) (*Metric, error) {
	return &Metric{
		Name:       name,
		Type:       vo.GaugeMetricType.String(),
		GaugeValue: &value,
	}, nil
}

func (m Metric) MapToDTO() dto.MetricDTO {
	return dto.MetricDTO{
		Type:         m.Type,
		Name:         m.Name,
		CounterValue: m.CounterValue,
		GaugeValue:   m.GaugeValue,
	}
}
