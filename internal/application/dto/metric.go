package dto

import (
	"strconv"

	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
)

type MetricDTO struct {
	Type         string   `json:"type"`
	Name         string   `json:"name"`
	CounterValue *int64   `json:"counter_value,omitempty"`
	GaugeValue   *float64 `json:"gauge_value,omitempty"`
}

func (m MetricDTO) Value() any {
	if m.CounterValue != nil {
		return *m.CounterValue
	}
	if m.GaugeValue != nil {
		return *m.GaugeValue
	}
	return "-"
}

func (dto *MetricDTO) IsCounter() bool {
	return dto.Type == vo.CounterMetricType.String()
}
func (dto *MetricDTO) IsGauge() bool {
	return dto.Type == vo.GaugeMetricType.String()
}

func NewMetric(metricType, metricName, metricValue string) (*MetricDTO, error) {
	dto := &MetricDTO{
		Type: metricType,
		Name: metricName,
	}

	switch metricType {
	case vo.CounterMetricType.String():
		v, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return &MetricDTO{}, errors.ErrHTTPStatusBadRequest
		}
		dto.CounterValue = &v
	case vo.GaugeMetricType.String():
		v, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return &MetricDTO{}, errors.ErrHTTPStatusBadRequest
		}
		dto.GaugeValue = &v
	}

	return dto, nil
}
