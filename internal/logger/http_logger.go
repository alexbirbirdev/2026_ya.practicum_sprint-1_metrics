package logger

import (
	"net/http"
	"time"
)

type responseData struct {
	status int
	size   int
}

type logResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (l *logResponseWriter) Write(b []byte) (int, error) {
	size, err := l.ResponseWriter.Write(b)
	l.responseData.size += size
	return size, err
}

func (l *logResponseWriter) WriteHeader(statusCode int) {
	l.ResponseWriter.WriteHeader(statusCode)
	l.responseData.status = statusCode
}

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := logResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		h.ServeHTTP(&lw, r)

		uri := r.URL
		method := r.Method

		duration := time.Since(start)

		AppLogger.Sugar().Info(
			method, uri, duration, responseData.status, " size: ", responseData.size,
		)
	}
	return http.HandlerFunc(logFn)
}
