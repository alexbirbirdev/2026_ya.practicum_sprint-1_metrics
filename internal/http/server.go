package http

import (
	"net/http"
)

func NewHTTPServer(r *HTTPRouter) {
	http.ListenAndServe(":8080", r.Chi)
}
