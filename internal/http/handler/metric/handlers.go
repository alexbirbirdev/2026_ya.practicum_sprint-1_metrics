package metric

import (
	"bytes"
	"encoding/json"
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
	IncrementCounter(in dto.Metrics) (dto.Metrics, error)
	UpdateGauge(in dto.Metrics) (dto.Metrics, error)
	GetAllMetrics() ([]dto.Metrics, error)
	GetMetricByNameAndType(metricName string, metricType vo.MetricType) (dto.Metrics, error)
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
		_, err := h.metricsUsecases.IncrementCounter(*request)
		if err != nil {
			log.Printf("increment counter: %v", err)
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusOK)
		return
	}
	_, err = h.metricsUsecases.UpdateGauge(*request)
	if err != nil {
		log.Printf("update gauge: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h MetricHandlers) UpdateMetric(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != vo.ContentTypeApplicationJSON {
		log.Printf("wrong content-type")
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var requestDTO dto.Metrics
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &requestDTO); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := requestDTO.Validate(); err != nil {
		log.Printf("validation error: %v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if requestDTO.IsCounter() {
		updatedMetric, err := h.metricsUsecases.IncrementCounter(requestDTO)
		if err != nil {
			log.Printf("increment counter: %v", err)
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		out, err := json.Marshal(updatedMetric)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Header().Add("Content-type", vo.ContentTypeApplicationJSON)
		res.WriteHeader(http.StatusOK)
		res.Write(out)
		return
	}
	updatedMetric, err := h.metricsUsecases.UpdateGauge(requestDTO)
	if err != nil {
		log.Printf("update gauge: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	out, err := json.Marshal(updatedMetric)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-type", vo.ContentTypeApplicationJSON)
	res.WriteHeader(http.StatusOK)
	res.Write(out)
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
	case metric.Delta != nil:
		value = fmt.Sprintf("%d", *metric.Delta)
	case metric.Value != nil:
		value = fmt.Sprintf("%g", *metric.Value)
	default:
		http.Error(res, "no value", http.StatusNotFound)
		return
	}
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(value))
}

func (h MetricHandlers) GetMetricValue(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != vo.ContentTypeApplicationJSON {
		log.Printf("wrong content-type")
		http.Error(res, "wrong content-type", http.StatusBadRequest)
		return
	}
	var requestDTO dto.Metrics
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		log.Printf("error read body")
		http.Error(res, "failed to read body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(buf.Bytes(), &requestDTO); err != nil {
		http.Error(res, "failed to unmarshal", http.StatusBadRequest)
		return
	}
	if !vo.MetricType(requestDTO.MType).Validate() {
		http.Error(res, "wrong metric type", http.StatusBadRequest)
		return
	}
	metric, err := h.metricsUsecases.GetMetricByNameAndType(requestDTO.ID, vo.MetricType(requestDTO.MType))
	if err != nil {
		if errors.Is(err, appErrors.ErrNotFound) {
			http.Error(res, "no metric", http.StatusNotFound)
			return
		}
		log.Printf("get metric by name and type: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(metric)
	if err != nil {
		http.Error(res, "failed to marshal", http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-type", vo.ContentTypeApplicationJSON)
	res.WriteHeader(http.StatusOK)
	res.Write(out)
}
