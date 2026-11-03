package dto

import (
	"fmt"
	"strconv"

	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
)

type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func (dto *Metrics) IsCounter() bool {
	return dto.MType == vo.CounterMetricType.String()
}
func (dto *Metrics) IsGauge() bool {
	return dto.MType == vo.GaugeMetricType.String()
}
func (dto Metrics) GetValue() any {
	if dto.Delta != nil {
		return *dto.Delta
	}
	if dto.Value != nil {
		return *dto.Value
	}
	return "-"
}

func (dto *Metrics) Validate() error {
	mt := vo.MetricType(dto.MType)
	if !mt.Validate() {
		return fmt.Errorf("%w: %q", errors.ErrUnknownType, dto.MType)
	}

	switch mt {
	case vo.CounterMetricType:
		if dto.Delta == nil {
			return errors.ErrCounterNoDelta
		}
		if dto.Value != nil {
			return errors.ErrCounterHasValue
		}
	case vo.GaugeMetricType:
		if dto.Value == nil {
			return errors.ErrGaugeNoValue
		}
		if dto.Delta != nil {
			return errors.ErrGaugeHasDelta
		}
	default:
		return fmt.Errorf("%w: %q", errors.ErrUnknownType, dto.MType)
	}

	return nil
}

func NewMetric(metricType, metricName, metricValue string) (*Metrics, error) {
	if !vo.MetricType(metricType).Validate() {
		return &Metrics{}, errors.ErrHTTPStatusBadRequest
	}

	dto := &Metrics{
		MType: metricType,
		ID:    metricName,
	}

	switch metricType {
	case vo.CounterMetricType.String():
		v, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return &Metrics{}, errors.ErrHTTPStatusBadRequest
		}
		dto.Delta = &v
	case vo.GaugeMetricType.String():
		v, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return &Metrics{}, errors.ErrHTTPStatusBadRequest
		}
		dto.Value = &v
	}

	return dto, nil
}
