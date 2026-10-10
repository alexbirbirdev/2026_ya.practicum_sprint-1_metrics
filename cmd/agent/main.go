package main

import (
	"flag"
	"time"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/agent"
	handler "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

func main() {
	addr := flag.String("a", "localhost:8080", "адрес эндопинта HTTP сервера")
	reportSec := flag.Int("r", 10, "частота отправки метрик на сервер")
	pollSec := flag.Int("p", 2, "частота опроса метрик")
	flag.Parse()

	pollInterval := time.Duration(*pollSec) * time.Second
	reportInterval := time.Duration(*reportSec) * time.Second

	metricKeys := db.NewMetricKeys()
	metricValues := db.NewMetricsValues()

	baseURL := "http://" + *addr
	handler := handler.NewHandler(baseURL)

	usecases := usecase.NewAgentUsecase(*handler)

	go usecases.CollectMetrics(*metricKeys, *metricValues, pollInterval)

	usecases.SendMetrics(*metricKeys, *metricValues, reportInterval)
}
