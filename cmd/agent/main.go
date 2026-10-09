package main

import (
	"time"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/agent"
	handler "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

func main() {
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second

	metricKeys := db.NewMetricKeys()
	metricValues := db.NewMetricsValues()

	baseURL := "http://localhost:8080"
	handler := handler.NewHandler(baseURL)

	usecases := usecase.NewAgentUsecase(*handler)

	go usecases.CollectMetrics(*metricKeys, *metricValues, pollInterval)

	usecases.SendMetrics(*metricKeys, *metricValues, reportInterval)
}
