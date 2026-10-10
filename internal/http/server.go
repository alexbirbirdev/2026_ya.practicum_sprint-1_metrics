package http

import (
	"net/http"
)

func NewHTTPServer(addr string, r *HTTPRouter) {
	http.ListenAndServe(addr, r.Chi)
}
