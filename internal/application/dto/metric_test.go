package dto_test

import (
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

func TestMetricDTO_Value(t *testing.T) {
	tests := []struct {
		name        string
		metricType  string
		metricName  string
		metricValue string
		want        any
	}{
		{
			name:        "valid counter value",
			metricType:  "counter",
			metricValue: "123",
			want:        int64(123),
		},
		{
			name:        "null counter ",
			metricType:  "counter",
			metricValue: "0",
			want:        int64(0),
		},
		{
			name:        "valid gauge value",
			metricType:  "gauge",
			metricValue: "123.123",
			want:        float64(123.123),
		},
		{
			name:        "null gauge",
			metricType:  "gauge",
			metricValue: "0",
			want:        float64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := dto.NewMetric(tt.metricType, tt.metricName, tt.metricValue)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := m.Value()
			if got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMetric(t *testing.T) {
	tests := []struct {
		name        string
		metricType  string
		metricName  string
		metricValue string
		want        *dto.MetricDTO
		wantErr     bool
	}{
		{
			name:        "valid counter metric",
			metricType:  "counter",
			metricName:  "test name",
			metricValue: "1234523",
			want: &dto.MetricDTO{
				Type:         vo.CounterMetricType.String(),
				Name:         "test name",
				CounterValue: ptr(int64(1234523)),
			},
			wantErr: false,
		},
		{
			name:        "valid gauge metric",
			metricType:  "gauge",
			metricName:  "test name",
			metricValue: "1234523.34",
			want: &dto.MetricDTO{
				Type:       vo.GaugeMetricType.String(),
				Name:       "test name",
				GaugeValue: ptr(1234523.34),
			},
			wantErr: false,
		},
		{
			name:        "error metric type",
			metricType:  "counter1",
			metricName:  "test name",
			metricValue: "1234523",
			want:        &dto.MetricDTO{},
			wantErr:     false,
		},
		{
			name:        "wrong counter value",
			metricType:  "counter",
			metricName:  "test name",
			metricValue: "123.123",
			want:        &dto.MetricDTO{},
			wantErr:     true,
		},
		{
			name:        "error wrong value type",
			metricType:  "counter",
			metricName:  "test name",
			metricValue: "msekmde",
			want:        &dto.MetricDTO{},
			wantErr:     true,
		},
		{
			name:        "error counter null",
			metricType:  "counter",
			metricName:  "test name",
			metricValue: "",
			want:        &dto.MetricDTO{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := dto.NewMetric(tt.metricType, tt.metricName, tt.metricValue)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewMetric() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewMetric() succeeded unexpectedly")
			}
			if got == tt.want {
				t.Errorf("NewMetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

func TestMetricDTO_IsCounter(t *testing.T) {
	tests := []struct {
		name   string
		metric dto.MetricDTO
		want   bool
	}{
		{
			name: "true",
			metric: dto.MetricDTO{
				Type:         vo.CounterMetricType.String(),
				Name:         "random",
				CounterValue: ptr(int64(1231)),
			},
			want: true,
		},
		{
			name: "false",
			metric: dto.MetricDTO{
				Type:         vo.GaugeMetricType.String(),
				Name:         "random",
				CounterValue: ptr(int64(1231)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.metric.IsCounter()
			if got != tt.want {
				t.Errorf("IsCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetricDTO_IsGauge(t *testing.T) {
	tests := []struct {
		name   string
		metric dto.MetricDTO
		want   bool
	}{
		{
			name: "false",
			metric: dto.MetricDTO{
				Type:         vo.CounterMetricType.String(),
				Name:         "random",
				CounterValue: ptr(int64(1231)),
			},
			want: false,
		},
		{
			name: "false",
			metric: dto.MetricDTO{
				Type:         vo.GaugeMetricType.String(),
				Name:         "random",
				CounterValue: ptr(int64(1231)),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.metric.IsGauge()
			if got != tt.want {
				t.Errorf("IsGauge() = %v, want %v", got, tt.want)
			}
		})
	}
}
