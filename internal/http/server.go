package http

import (
	"log"
	"net/http"
)

func NewHTTPServer(addr string, r *HTTPRouter) {
	if err := http.ListenAndServe(addr, r.Chi); err != nil {
		log.Fatalf("http server failed: %v", err)
	}
}
