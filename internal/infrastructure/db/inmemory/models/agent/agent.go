package agent

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
	CounterMetricsValues map[string]int64
	GaugeMetricsValues   map[string]float64
}

func NewMetricsValues() *MetricsValues {
	return &MetricsValues{
		CounterMetricsValues: make(map[string]int64),
		GaugeMetricsValues:   make(map[string]float64),
	}
}
