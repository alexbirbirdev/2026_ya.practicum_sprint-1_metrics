package models

import "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"

type InMemoryMetrics struct {
	Metrics []Metric
}

type Metric struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Hash  string   `json:"hash,omitempty"`
}

func NewInMemoryMetrics() (*InMemoryMetrics, error) {
	return &InMemoryMetrics{
		Metrics: make([]Metric, 0),
	}, nil
}

func (m *Metric) MapToEntity() *entity.Metric {
	var delta int
	if m.Delta != nil {
		delta = int(*m.Delta)
	}
	var value float64
	if m.Value != nil {
		value = *m.Value
	}
	return &entity.Metric{
		Name:         m.ID,
		Type:         m.MType,
		CounterValue: &delta,
		GaugeValue:   &value,
	}
}

func MapToModel(m *entity.Metric) *Metric {
	var delta int64
	if m.CounterValue != nil {
		delta = int64(*m.CounterValue)
	}
	var value float64
	if m.GaugeValue != nil {
		value = *m.GaugeValue
	}
	return &Metric{
		ID:    m.Name,
		MType: m.Type,
		Delta: &delta,
		Value: &value,
		Hash:  "default",
	}
}

// из описания задания мне не было понятно, для чего хеш, но пока оставил на будущее
