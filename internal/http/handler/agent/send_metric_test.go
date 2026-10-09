package agent_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	usecase "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/usecase/metric"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	mocks "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/interfaces/mocks"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	httplib "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	metricHandlers "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/metric"
	"github.com/go-openapi/testify/v2/assert"
)

func TestHandler_SendMetric(t *testing.T) {
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
		in      dto.MetricDTO
		want    int
		wantErr bool
	}{
		{
			in: dto.MetricDTO{
				Type:         vo.CounterMetricType.String(),
				Name:         "test",
				CounterValue: ptr(int64(123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.MetricDTO{
				Type:       vo.GaugeMetricType.String(),
				Name:       "test",
				GaugeValue: ptr(float64(123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.MetricDTO{
				Type:       vo.GaugeMetricType.String(),
				Name:       "test",
				GaugeValue: ptr(float64(123.123)),
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			in: dto.MetricDTO{
				Type:       "anoter",
				Name:       "test123",
				GaugeValue: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
		{
			in: dto.MetricDTO{
				Type:       "anoter sdfsd f sdf ",
				Name:       "test123",
				GaugeValue: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
		{
			in: dto.MetricDTO{
				Type:       " ",
				Name:       " ",
				GaugeValue: ptr(float64(123)),
			},
			want:    http.StatusBadRequest,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in.Name, func(t *testing.T) {
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
