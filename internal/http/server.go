package http

import (
	"log"
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/logger"
)

func NewHTTPServer(addr string, r *HTTPRouter) {
	routerWithLogs := logger.WithLogging(r.Chi)
	if err := http.ListenAndServe(addr, routerWithLogs); err != nil {
		log.Fatalf("http server failed: %v", err)
	}
}
