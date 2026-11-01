package compress

import (
	"net/http"
	"strings"
)

func WithCompression(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") == "gzip" {
			cr, err := newCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			r.Header.Del("Content-Encoding")
			r.Header.Del("Content-Length")
		}

		originalWriter := http.ResponseWriter(w)
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			compressWriter := newCompressWriter(w)
			originalWriter = compressWriter
			defer compressWriter.Close()
		}

		h.ServeHTTP(originalWriter, r)
	})
}
