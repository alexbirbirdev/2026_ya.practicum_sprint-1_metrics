package main

import (
	"fmt"
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler"
	models "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/repository"
)

func main() {
	inMemoryDB, err := models.NewInMemoryMetrics()
	if err != nil {
		panic(fmt.Errorf("fail to initialize DB"))
	}

	metricsRepository := repository.NewMetricsRepository(inMemoryDB)

	metricsUsecase := usecase.NewMetricUsecases(metricsRepository)

	metricsHandler := handler.NewMetricHandler(*metricsUsecase)

	http.HandleFunc(`/update/`, metricsHandler.AcceptMetric)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
