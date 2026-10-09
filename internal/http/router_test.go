package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	mocks "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/interfaces/mocks"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	httplib "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestHTTPRouter(t *testing.T) {
	metricsRepository := &mocks.MockMetricsRepository{
		ListMetricsFn: func() ([]entity.Metric, error) {
			counter := int64(42)
			return []entity.Metric{
				{
					Name:         "requests",
					Type:         vo.CounterMetricType.String(),
					CounterValue: &counter,
				},
			}, nil
		},
	}

	metricsUsecases := usecase.NewMetricUsecases(metricsRepository)

	metricsHandlers := metricHandlers.NewMetricHandler(*metricsUsecases)

	handlers := httplib.HTTPHandlers{
		Metrics: metricsHandlers,
	}

	router := httplib.NewHTTPRouter(handlers)
	ts := httptest.NewServer(router.Chi)
	defer ts.Close()

	tests := []struct {
		url         string
		method      string
		contentType string
		want        int
	}{
		{
			url:         "/",
			method:      http.MethodGet,
			contentType: "application/json",
			want:        http.StatusOK,
		},
		{
			url:         "/",
			method:      http.MethodPost,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodPut,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodDelete,
			contentType: "application/json",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodPost,
			contentType: "application/json",
			want:        http.StatusUnsupportedMediaType,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodPut,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/some/123",
			method:      http.MethodDelete,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
		{
			url:         "/update/counter/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/update/counter1/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/update/counter/some/123.123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusBadRequest,
		},
		{
			url:         "/update/gauge/some/123.123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/update/gauge/some/123.123sdf",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusBadRequest,
		},
		{
			url:         "/update/gauge/some/123",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusOK,
		},
		{
			url:         "/value/gauge/some",
			method:      http.MethodGet,
			contentType: "text/plain",
			want:        http.StatusNotFound,
		},
		{
			url:         "/value/gauge/some",
			method:      http.MethodPost,
			contentType: "text/plain",
			want:        http.StatusMethodNotAllowed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			url := fmt.Sprintf("%s%s", ts.URL, tt.url)
			req, err := http.NewRequest(tt.method, url, nil)
			req.Header.Set("Content-Type", tt.contentType)
			require.NoError(t, err)
			w := httptest.NewRecorder()

			router.Chi.ServeHTTP(w, req)

			assert.Equal(t, tt.want, w.Code)
		})
	}
}
