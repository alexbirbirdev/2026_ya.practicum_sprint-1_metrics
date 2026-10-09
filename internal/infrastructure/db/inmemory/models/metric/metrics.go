package models

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
)

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
	metric := &entity.Metric{
		Name:         m.ID,
		Type:         m.MType,
		CounterValue: m.Delta,
		GaugeValue:   m.Value,
	}
	return metric
}

func MapToModel(m *entity.Metric) *Metric {
	return &Metric{
		ID:    m.Name,
		MType: m.Type,
		Delta: m.CounterValue,
		Value: m.GaugeValue,
		Hash:  "default",
	}
}

// из описания задания мне не было понятно, для чего хеш, но пока оставил на будущее
