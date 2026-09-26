package handler

import (
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase"
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
	if req.Method != http.MethodPost {
		http.Error(res, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, "Wrong Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	request, err := dto.NewUpdateMetricRequest(req.URL.Path)
	if err != nil {
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
