package dto

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type UpdateMetricsRequest struct {
	Type         string
	Name         string
	CounterValue *int
	GaugeValue   *float64
}

func (dto *UpdateMetricsRequest) IsCounter() bool {
	if dto.Type == vo.Counter {
		return true
	}
	return false
}
func (dto *UpdateMetricsRequest) IsGauge() bool {
	if dto.Type == vo.Gauge {
		return true
	}
	return false
}

func NewUpdateMetricRequest(URL string) (*UpdateMetricsRequest, error) {
	path := strings.Split(strings.Trim(URL, "/"), "/")
	if len(path) != 4 {
		return &UpdateMetricsRequest{}, fmt.Errorf(strconv.Itoa(http.StatusBadRequest))
	}

	metricType := path[1]
	if metricType != vo.Counter && metricType != vo.Gauge {
		return &UpdateMetricsRequest{}, fmt.Errorf(strconv.Itoa(http.StatusBadRequest))
	}

	metricName := path[2]
	metricValue := path[3]

	var metricCounterValue int
	var metricGaugeValue float64
	var err error

	if metricType == vo.Counter {
		metricCounterValue, err = strconv.Atoi(metricValue)
		if err != nil {
			return &UpdateMetricsRequest{}, fmt.Errorf(strconv.Itoa(http.StatusBadRequest))
		}
	}
	if metricType == vo.Gauge {
		metricGaugeValue, err = strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return &UpdateMetricsRequest{}, fmt.Errorf(strconv.Itoa(http.StatusBadRequest))
		}
	}

	return &UpdateMetricsRequest{
		Type:         metricType,
		Name:         metricName,
		CounterValue: &metricCounterValue,
		GaugeValue:   &metricGaugeValue,
	}, nil
}
