package valueobject_test

import (
	"testing"

	valueobject "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	"github.com/go-openapi/testify/v2/assert"
)

func TestMetricType_Validate(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "counter",
			want: true,
		},
		{
			name: "gauge",
			want: true,
		},
		{
			name: "something",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := valueobject.MetricType(tt.name)
			got := m.Validate()

			assert.Equal(t, got, tt.want)
		})
	}
}
