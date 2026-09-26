package dto

import (
	"strconv"
	"strings"

	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
)

type UpdateMetricsRequest struct {
	Type         string
	Name         string
	CounterValue *int
	GaugeValue   *float64
}

func (dto *UpdateMetricsRequest) IsCounter() bool {
	return dto.Type == vo.Counter
}
func (dto *UpdateMetricsRequest) IsGauge() bool {
	return dto.Type == vo.Gauge
}

func NewUpdateMetricRequest(URL string) (*UpdateMetricsRequest, error) {
	path := strings.Split(strings.Trim(URL, "/"), "/")
	if len(path) != 4 {
		return &UpdateMetricsRequest{}, errors.ErrHTTPStatusNotFound
	}

	metricType := path[1]
	if metricType != vo.Counter && metricType != vo.Gauge {
		return &UpdateMetricsRequest{}, errors.ErrHTTPStatusBadRequest
	}

	metricName := path[2]
	metricValue := path[3]

	var metricCounterValue int
	var metricGaugeValue float64
	var err error

	if metricType == vo.Counter {
		metricCounterValue, err = strconv.Atoi(metricValue)
		if err != nil {
			return &UpdateMetricsRequest{}, errors.ErrHTTPStatusBadRequest
		}
	}
	if metricType == vo.Gauge {
		metricGaugeValue, err = strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return &UpdateMetricsRequest{}, errors.ErrHTTPStatusBadRequest
		}
	}

	return &UpdateMetricsRequest{
		Type:         metricType,
		Name:         metricName,
		CounterValue: &metricCounterValue,
		GaugeValue:   &metricGaugeValue,
	}, nil
}
