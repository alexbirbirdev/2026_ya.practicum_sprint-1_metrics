package http

import (
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPHandlers struct {
	Metrics *metric.MetricHandlers
	Server  *server.ServerHandlers
}

type HTTPRouter struct {
	Chi chi.Router
}

func NewHTTPRouter(cfg HTTPHandlers) *HTTPRouter {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)

	r.Get("/ping", cfg.Server.Ping)

	r.Get("/", cfg.Metrics.GetAllMetrics)
	r.Get("/value/{metricType}/{metricName}", cfg.Metrics.GetMetricByName)
	r.Post("/value", cfg.Metrics.GetMetricValue)
	r.Post("/update/{metricType}/{metricName}/{metricValue}", cfg.Metrics.AcceptMetric)
	r.Post("/update", cfg.Metrics.UpdateMetric)

	return &HTTPRouter{Chi: r}
}
