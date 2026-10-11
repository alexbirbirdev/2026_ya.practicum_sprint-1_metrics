package agent

import "sync"

type MetricsKeys struct {
	CounterMetricsKeys []string
	GaugeMetricsKeys   []string
}

func NewMetricKeys() *MetricsKeys {
	return &MetricsKeys{
		CounterMetricsKeys: []string{"PollCount"},
		GaugeMetricsKeys: []string{
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
			"RandomValue",
		},
	}
}

type MetricsValues struct {
	mu                  sync.RWMutex
	counterMetricValues map[string]int64
	gaugeMetricValues   map[string]float64
}

func NewMetricsValues() *MetricsValues {
	return &MetricsValues{
		counterMetricValues: make(map[string]int64),
		gaugeMetricValues:   make(map[string]float64),
	}
}

func (v *MetricsValues) IncCounter(name string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.counterMetricValues[name]++
}
func (v *MetricsValues) SetCounter(name string, value int64) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.counterMetricValues[name] = value
}

func (v *MetricsValues) SetGauge(name string, value float64) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.gaugeMetricValues[name] = value
}

func (v *MetricsValues) Counter(name string) int64 {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.counterMetricValues[name]
}

func (v *MetricsValues) Gauge(name string) float64 {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.gaugeMetricValues[name]
}
