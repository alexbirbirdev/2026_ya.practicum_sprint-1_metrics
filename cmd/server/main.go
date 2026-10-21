package main

import (
	"log"

	config "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/config/server"
	http "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/repository"
)

func main() {
	cfg := config.GetConfig()

	inMemoryDB, err := models.NewInMemoryMetrics()
	if err != nil {
		log.Fatalf("fail to initialize DB: %v", err)
	}

	metricsRepository := repository.NewMetricsRepository(inMemoryDB)

	metricsUsecases := usecase.NewMetricUsecases(metricsRepository)

	metricsHandlers := metricHandlers.NewMetricHandler(metricsUsecases)

	handlers := http.HTTPHandlers{
		Metrics: metricsHandlers,
	}

	router := http.NewHTTPRouter(handlers)

	http.NewHTTPServer(cfg.HTTP.Addr, router)
}
