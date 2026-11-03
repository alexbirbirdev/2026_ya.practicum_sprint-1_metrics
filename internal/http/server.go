package http

import (
	"log"
	"net/http"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/middleware/compress"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/logger"
)

func NewHTTPServer(addr string, r *HTTPRouter) {
	routerWithCompression := compress.WithCompression(r.Chi)
	routerWithLogs := logger.WithLogging(routerWithCompression)
	if err := http.ListenAndServe(addr, routerWithLogs); err != nil {
		log.Fatalf("http server failed: %v", err)
	}
}
