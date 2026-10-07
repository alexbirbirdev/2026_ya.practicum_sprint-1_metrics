package agent

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	service "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/service/agent"
	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
	db "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/infrastructure/db/inmemory/models/agent"
)

type AgentUsecase struct{}

func NewAgentUsecase() *AgentUsecase {
	return &AgentUsecase{}
}

func (u *AgentUsecase) CollectMetrics(keys db.MetricsKeys, values db.MetricsValues, interval time.Duration) {
	for {
		for _, m := range keys.CounterMetricsKeys {
			if m == "PollCount" {
				values.CounterMetricsValues[m] += 1
			}
		}
		for _, m := range keys.GaugeMetricsKeys {
			if m == "RandomValue" {
				values.GaugeMetricsValues[m] = rand.Float64()
			} else {
				values.GaugeMetricsValues[m] = service.GetMetricValue(m)
			}
		}
		time.Sleep(interval)
	}
}
func (u *AgentUsecase) SendMetrics(keys db.MetricsKeys, values db.MetricsValues, interval time.Duration) {
	for {
		for _, m := range keys.CounterMetricsKeys {
			if err := sendMetric(vo.Counter, m, values.CounterMetricsValues[m]); err != nil {
				panic(err)
			}
		}
		for _, m := range keys.GaugeMetricsKeys {
			if err := sendMetric(vo.Gauge, m, values.GaugeMetricsValues[m]); err != nil {
				panic(err)
			}
		}
		time.Sleep(interval)
	}
}

var client = &http.Client{
	Timeout: 1 * time.Second,
}

func sendMetric(metricType, metricName string, metricValue any) error {

	url := fmt.Sprintf("http://localhost:8080/update/%s/%s/%v", metricType, metricName, metricValue)
	response, err := client.Post(url, "text/plain", nil)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not 200")
	}
	response.Body.Close()
	return nil
}
