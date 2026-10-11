package main

import (
	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/agent"
	config "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/config/agent"
	handler "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

func main() {
	cfg := config.ParseAgent()
	metricKeys := db.NewMetricKeys()
	metricValues := db.NewMetricsValues()

	handler := handler.NewHandler(cfg.HTTP.BaseURL())

	usecases := usecase.NewAgentUsecase(*handler)

	go usecases.CollectMetrics(metricKeys, metricValues, cfg.Collector.PollInterval)

	usecases.SendMetrics(metricKeys, metricValues, cfg.Reporter.ReportInterval)
}
