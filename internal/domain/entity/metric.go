package entity

import (
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type Metric struct {
	Name         string
	Type         string
	CounterValue *int
	GaugeValue   *float64
}

func (m *Metric) IncrementCounterValue(value int) error {
	if m.CounterValue != nil {
		*m.CounterValue += value
	}
	return nil
}

func (m *Metric) UpdateGaugeValue(value float64) error {
	if m.GaugeValue != nil {
		*m.GaugeValue = value
	}
	return nil
}

func NewCounterMetric(
	name string,
	value int,
) (*Metric, error) {
	return &Metric{
		Name:         name,
		Type:         vo.Counter,
		CounterValue: &value,
	}, nil
}
func NewGaugeMetric(
	name string,
	value float64,
) (*Metric, error) {
	return &Metric{
		Name:       name,
		Type:       vo.Gauge,
		GaugeValue: &value,
	}, nil
}
