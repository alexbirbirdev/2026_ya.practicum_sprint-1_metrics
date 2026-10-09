package metric

import (
	"fmt"
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	"github.com/go-chi/chi/v5"
)

type MetricHandlers struct {
	metricsUsecases usecase.MetricUsecases
}

func NewMetricHandler(
	metricsUsecases usecase.MetricUsecases,
) *MetricHandlers {
	return &MetricHandlers{
		metricsUsecases: metricsUsecases,
	}
}

func (h MetricHandlers) AcceptMetric(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, "wrong content-Type", http.StatusUnsupportedMediaType)
		return
	}
	metricType := chi.URLParam(req, "metricType")
	if !vo.MetricType(metricType).Validate() {
		http.Error(res, "wrong metric type", http.StatusBadRequest)
		return
	}
	metricName := chi.URLParam(req, "metricName")
	metricValue := chi.URLParam(req, "metricValue")

	request, err := dto.NewMetric(metricType, metricName, metricValue)
	if err != nil {
		if err == errors.ErrNotFound {
			http.Error(res, "Wrong Request", http.StatusNotFound)
			return
		} else if err == errors.ErrHTTPStatusBadRequest {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}
		http.Error(res, "Validation Error", http.StatusBadRequest)
		return
	}
	if request.IsCounter() {
		err := h.metricsUsecases.IncrementCounter(*request)
		if err != nil {
			http.Error(res, "Server error", http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusOK)
		return
	}
	err = h.metricsUsecases.UpdateGauge(*request)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h MetricHandlers) GetAllMetrics(res http.ResponseWriter, req *http.Request) {
	list, err := h.metricsUsecases.GetAllMetrics()
	if err != nil {
		http.Error(res, "server error", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	if err := metricsListHTML.ExecuteTemplate(res, "metric.html", list); err != nil {
		http.Error(res, "server error", http.StatusInternalServerError)
		return
	}
}

func (h MetricHandlers) GetMetricByName(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "metricType")
	if !vo.MetricType(metricType).Validate() {
		http.Error(res, "wrong metric type", http.StatusBadRequest)
		return
	}
	metricName := chi.URLParam(req, "metricName")

	metric, err := h.metricsUsecases.GetMetricByNameAndType(metricName, vo.MetricType(metricType))
	if err != nil {
		if err == errors.ErrNotFound {
			http.Error(res, "no metric", http.StatusNotFound)
			return
		}
		http.Error(res, "something wrong", http.StatusInternalServerError)
		return
	}
	var value string
	switch {
	case metric.CounterValue != nil:
		value = fmt.Sprintf("%d", *metric.CounterValue)
	case metric.GaugeValue != nil:
		value = fmt.Sprintf("%g", *metric.GaugeValue)
	default:
		http.Error(res, "no value", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(value))
}
