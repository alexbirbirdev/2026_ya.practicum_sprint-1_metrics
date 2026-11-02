package main

import (
	"log"

	config "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/config/server"
	http "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/logger"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	filestorage "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/file_storage"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/repository"
)

func main() {
	cfg := config.GetConfig()

	logger.InitAppLogger("info")

	inMemoryDB, err := models.NewInMemoryMetrics()
	if err != nil {
		log.Fatalf("fail to initialize DB: %v", err)
	}

	fileStorage := filestorage.FileStorage{
		Path: cfg.FileStorage.FileStoragePath,
	}

	baseRepository := repository.NewMetricsRepository(inMemoryDB)

	if cfg.FileStorage.Restore {
		loaded, err := fileStorage.Load()
		if err != nil {
			log.Fatalf("failed to load metrics from file: %v", err)
		}
		if len(loaded) > 0 {
			baseRepository.ReplaceAllMetrics(loaded)
		}
	}

	metricsRepository := repository.NewPersistingMetricsRepository(
		baseRepository,
		fileStorage,
		cfg.FileStorage.StoreInterval,
	)

	metricsRepository.RunPeriodicMetricsSaving()

	metricsUsecases := usecase.NewMetricUsecases(metricsRepository)

	metricsHandlers := metricHandlers.NewMetricHandler(metricsUsecases)

	handlers := http.HTTPHandlers{
		Metrics: metricsHandlers,
	}

	router := http.NewHTTPRouter(handlers)

	http.NewHTTPServer(cfg.HTTP.Addr, router)
}
