package entity_test

import (
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/go-openapi/testify/v2/assert"
)

func TestNewCounterMetric(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		value    int64
		want     *entity.Metric
		wantErr  bool
	}{
		{
			testName: "valid",
			name:     "random",
			value:    int64(123),
			want: &entity.Metric{
				Name:         "random",
				Type:         vo.CounterMetricType.String(),
				CounterValue: ptr(int64(123)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := entity.NewCounterMetric(tt.name, tt.value)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewCounterMetric() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewCounterMetric() succeeded unexpectedly")
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

func TestNewGaugeMetric(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		value    float64
		want     *entity.Metric
		wantErr  bool
	}{
		{
			testName: "valid",
			name:     "random",
			value:    float64(123.123123),
			want: &entity.Metric{
				Name:       "random",
				Type:       vo.GaugeMetricType.String(),
				GaugeValue: ptr(123.123123),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := entity.NewGaugeMetric(tt.name, tt.value)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewGaugeMetric() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewGaugeMetric() succeeded unexpectedly")
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

func ptr[T any](v T) *T { return &v }

func TestMetric_IncrementCounterValue(t *testing.T) {
	tests := []struct {
		name    string
		value   *int64
		source  entity.Metric
		wantSum int64
		wantErr bool
	}{
		{
			name:  "test",
			value: ptr(int64(111)),
			source: entity.Metric{
				Type:         vo.CounterMetricType.String(),
				Name:         "test",
				CounterValue: ptr(int64(111)),
			},
			wantSum: int64(222),
			wantErr: false,
		},
		{
			name:  "test 0",
			value: ptr(int64(0)),
			source: entity.Metric{
				Type:         vo.CounterMetricType.String(),
				Name:         "test",
				CounterValue: ptr(int64(0)),
			},
			wantSum: int64(0),
			wantErr: false,
		},
		{
			name:  "test 1",
			value: ptr(int64(0)),
			source: entity.Metric{
				Type:         vo.CounterMetricType.String(),
				Name:         "test",
				CounterValue: ptr(int64(1)),
			},
			wantSum: int64(1),
			wantErr: false,
		},
		{
			name:  "test 2",
			value: ptr(int64(1)),
			source: entity.Metric{
				Type:         vo.CounterMetricType.String(),
				Name:         "test",
				CounterValue: ptr(int64(0)),
			},
			wantSum: int64(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotErr := tt.source.IncrementCounterValue(tt.value)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("IncrementCounterValue() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("IncrementCounterValue() succeeded unexpectedly")
			}
			assert.Equal(t, tt.source.CounterValue, &tt.wantSum)
		})
	}
}

func TestMetric_UpdateGaugeValue(t *testing.T) {
	tests := []struct {
		name    string
		value   *float64
		source  entity.Metric
		wantRes float64
		wantErr bool
	}{
		{
			name:  "test",
			value: ptr(float64(111.123)),
			source: entity.Metric{
				Type:       vo.GaugeMetricType.String(),
				Name:       "test",
				GaugeValue: ptr(float64(111)),
			},
			wantRes: float64(111.123),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.source.UpdateGaugeValue(tt.value)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateGaugeValue() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateGaugeValue() succeeded unexpectedly")
			}
			assert.Equal(t, *tt.source.GaugeValue, tt.wantRes)
		})
	}
}
