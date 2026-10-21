package metric

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	appErrors "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	"github.com/go-chi/chi/v5"
)

type MetricUsecase interface {
	IncrementCounter(in dto.MetricDTO) error
	UpdateGauge(in dto.MetricDTO) error
	GetAllMetrics() ([]dto.MetricDTO, error)
	GetMetricByNameAndType(metricName string, metricType vo.MetricType) (dto.MetricDTO, error)
}

type MetricHandlers struct {
	metricsUsecases MetricUsecase
}

func NewMetricHandler(
	metricsUsecases MetricUsecase,
) *MetricHandlers {
	return &MetricHandlers{
		metricsUsecases: metricsUsecases,
	}
}

func (h MetricHandlers) AcceptMetric(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "metricType")
	if !vo.MetricType(metricType).Validate() {
		http.Error(res, "wrong metric type", http.StatusBadRequest)
		return
	}
	metricName := chi.URLParam(req, "metricName")
	metricValue := chi.URLParam(req, "metricValue")

	request, err := dto.NewMetric(metricType, metricName, metricValue)
	if err != nil {

		if errors.Is(err, appErrors.ErrNotFound) {
			http.Error(res, "Wrong Request", http.StatusNotFound)
			return
		} else if errors.Is(err, appErrors.ErrHTTPStatusBadRequest) {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			return
		}
		http.Error(res, "Validation Error", http.StatusBadRequest)
		return
	}
	if request.IsCounter() {
		err := h.metricsUsecases.IncrementCounter(*request)
		if err != nil {
			log.Printf("increment counter: %v", err)
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusOK)
		return
	}
	err = h.metricsUsecases.UpdateGauge(*request)
	if err != nil {
		log.Printf("update gauge: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h MetricHandlers) GetAllMetrics(res http.ResponseWriter, req *http.Request) {
	list, err := h.metricsUsecases.GetAllMetrics()
	if err != nil {
		log.Printf("list metrics: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	if err := metricsListHTML.ExecuteTemplate(res, "metric.html", list); err != nil {
		log.Printf("render metrics page: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		if errors.Is(err, appErrors.ErrNotFound) {
			http.Error(res, "no metric", http.StatusNotFound)
			return
		}
		log.Printf("get metric by name and type: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
