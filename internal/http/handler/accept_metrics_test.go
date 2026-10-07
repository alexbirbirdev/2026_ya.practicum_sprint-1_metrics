package handler_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/repository"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/go-openapi/testify/v2/require"
)

func TestMetricHandlers_AcceptMetric(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "wrong test",
			want: want{
				code:        415,
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inMemoryDB, err := models.NewInMemoryMetrics()
			if err != nil {
				panic(fmt.Errorf("fail to initialize DB"))
			}
			metricsRepository := repository.NewMetricsRepository(inMemoryDB)

			metricsUsecase := usecase.NewMetricUsecases(metricsRepository)

			metricsHandler := handler.NewMetricHandler(*metricsUsecase)

			request := httptest.NewRequest(
				"POST",
				"/update/gauge/name/123",
				nil,
			)

			w := httptest.NewRecorder()
			metricsHandler.AcceptMetric(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)
			defer result.Body.Close()
			require.NoError(t, err)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}
