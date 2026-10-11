package agent

import (
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/application/dto"
	service "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/service/agent"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	handler "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/http/handler/agent"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

type AgentUsecase struct {
	Handler handler.Handler
}

func NewAgentUsecase(
	handler handler.Handler,
) *AgentUsecase {
	return &AgentUsecase{
		Handler: handler,
	}
}

func (u *AgentUsecase) CollectMetrics(keys *db.MetricsKeys, values *db.MetricsValues, interval time.Duration) {
	for {
		for _, m := range keys.CounterMetricsKeys {
			if m == "PollCount" {
				values.IncCounter(m)
			}
		}
		for _, m := range keys.GaugeMetricsKeys {
			if m == "RandomValue" {
				values.SetGauge(m, rand.Float64())
			} else {
				values.SetGauge(m, service.GetMetricValue(m))
			}
		}
		time.Sleep(interval)
	}
}
func (u *AgentUsecase) SendMetrics(keys *db.MetricsKeys, values *db.MetricsValues, interval time.Duration) {
	for {
		for _, m := range keys.CounterMetricsKeys {
			value := values.Counter(m)
			metricDTO := dto.MetricDTO{
				Name:         m,
				Type:         vo.CounterMetricType.String(),
				CounterValue: &value,
			}
			status, err := u.Handler.SendMetric(metricDTO)
			if err != nil || status != http.StatusOK {
				break
			}
		}
		for _, m := range keys.GaugeMetricsKeys {
			value := values.Gauge(m)
			metricDTO := dto.MetricDTO{
				Name:       m,
				Type:       vo.GaugeMetricType.String(),
				GaugeValue: &value,
			}
			status, err := u.Handler.SendMetric(metricDTO)
			if err != nil || status != http.StatusOK {
				break
			}
		}
		values.SetCounter("PollCount", 0)
		time.Sleep(interval)
	}
}
