package agent_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	repomocks "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric/mocks"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	appErrors "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object/errors"
	httplib "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	"github.com/go-openapi/testify/v2/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_SendMetric(t *testing.T) {
	repo := repomocks.NewMockMetricsRepository(t)
	repo.EXPECT().FindMetricByNameAndType(mock.Anything, mock.Anything).
		Return(entity.Metric{}, appErrors.ErrNotFound)
	repo.EXPECT().CreateMetric(mock.Anything).Return(nil)
	repo.EXPECT().UpdateMetric(mock.Anything).Return(nil).Maybe()

	metricsUsecases := usecase.NewMetricUsecases(repo)

	metricsHandlers := metricHandlers.NewMetricHandler(metricsUsecases)

	handlers := httplib.HTTPHandlers{
		Metrics: metricsHandlers,
	}

	router := httplib.NewHTTPRouter(handlers)
	ts := httptest.NewServer(router.Chi)
	defer ts.Close()

	tests := []struct {
		in      dto.Metrics
		want    int
		wantErr bool
	}{
		{
			in: dto.Metrics{
				MType: vo.CounterMetricType.String(),
				ID:    "test",
				Delta: ptr(int64(123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.Metrics{
				MType: vo.GaugeMetricType.String(),
				ID:    "test",
				Value: ptr(float64(123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.Metrics{
				MType: vo.GaugeMetricType.String(),
				ID:    "test",
				Value: ptr(float64(123.123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.Metrics{
				MType: "anoter",
				ID:    "test123",
				Value: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
		{
			in: dto.Metrics{
				MType: "anoter sdfsd f sdf ",
				ID:    "test123",
				Value: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
		{
			in: dto.Metrics{
				MType: " ",
				ID:    " ",
				Value: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in.ID, func(t *testing.T) {
			h := agent.NewHandler(ts.URL)
			got, gotErr := h.SendMetric(tt.in)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SendMetric() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SendMetric() succeeded unexpectedly")
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
