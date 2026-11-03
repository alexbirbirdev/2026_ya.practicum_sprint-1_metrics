package filestorage

import (
	"encoding/json"
	"os"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/entity"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type FileStorage struct {
	Path string
}

func (fs FileStorage) Save(metrics []dto.Metrics) error {
	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.Path, data, 0666)
}

func (fs FileStorage) Load() ([]entity.Metric, error) {
	data, err := os.ReadFile(fs.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	var list []dto.Metrics
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	out := make([]entity.Metric, 0, len(list))
	for _, metric := range list {
		if err := metric.Validate(); err != nil {
			return nil, err
		}
		switch vo.MetricType(metric.MType) {
		case vo.GaugeMetricType:
			m, err := entity.NewGaugeMetric(metric.ID, *metric.Value)
			if err != nil {
				return nil, err
			}
			out = append(out, *m)
		case vo.CounterMetricType:
			m, err := entity.NewCounterMetric(metric.ID, *metric.Delta)
			if err != nil {
				return nil, err
			}
			out = append(out, *m)
		}
	}
	return out, nil
}
