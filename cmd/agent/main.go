package main

import (
	"time"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/agent"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

func main() {
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second

	metricKeys := db.NewMetricKeys()
	metricValues := db.NewMetricsValues()

	usecases := usecase.NewAgentUsecase()

	go usecases.CollectMetrics(*metricKeys, *metricValues, pollInterval)

	usecases.SendMetrics(*metricKeys, *metricValues, reportInterval)
}
