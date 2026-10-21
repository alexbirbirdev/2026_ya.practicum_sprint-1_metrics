package agent_test

import (
	"testing"

	"github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/service/agent"
)

func TestGetMetricValue_UnknownName(t *testing.T) {
	got := agent.GetMetricValue("definitely_unknown_metric")
	if got != 0 {
		t.Errorf("unknown metric should return 0, got %v", got)
	}
}

func TestGetMetricValue_KnownNames(t *testing.T) {

	names := []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			got := agent.GetMetricValue(name)
			if got < 0 {
				t.Errorf("GetMetricValue(%q) = %v, want >= 0", name, got)
			}
		})
	}
}
func TestGetMetricValue_GCCPUFraction(t *testing.T) {
	got := agent.GetMetricValue("GCCPUFraction")
	if got < 0 || got > 1 {
		t.Errorf("GCCPUFraction = %v, want in [0, 1]", got)
	}
}
