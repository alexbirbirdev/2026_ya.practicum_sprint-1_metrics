package agent

import "runtime"

func GetMetricValue(name string) float64 {
	var met runtime.MemStats
	runtime.ReadMemStats(&met)
	switch name {
	case "Alloc":
		return float64(met.Alloc)
	case "BuckHashSys":
		return float64(met.BuckHashSys)
	case "Frees":
		return float64(met.Frees)
	case "GCCPUFraction":
		return float64(met.GCCPUFraction)
	case "GCSys":
		return float64(met.GCSys)
	case "HeapAlloc":
		return float64(met.HeapAlloc)
	case "HeapIdle":
		return float64(met.HeapIdle)
	case "HeapInuse":
		return float64(met.HeapInuse)
	case "HeapObjects":
		return float64(met.HeapObjects)
	case "HeapReleased":
		return float64(met.HeapReleased)
	case "HeapSys":
		return float64(met.HeapSys)
	case "LastGC":
		return float64(met.LastGC)
	case "Lookups":
		return float64(met.Lookups)
	case "MCacheInuse":
		return float64(met.MCacheInuse)
	case "MCacheSys":
		return float64(met.MCacheSys)
	case "MSpanInuse":
		return float64(met.MSpanInuse)
	case "MSpanSys":
		return float64(met.MSpanSys)
	case "Mallocs":
		return float64(met.Mallocs)
	case "NextGC":
		return float64(met.NextGC)
	case "NumForcedGC":
		return float64(met.NumForcedGC)
	case "NumGC":
		return float64(met.NumGC)
	case "OtherSys":
		return float64(met.OtherSys)
	case "PauseTotalNs":
		return float64(met.PauseTotalNs)
	case "StackInuse":
		return float64(met.StackInuse)
	case "StackSys":
		return float64(met.StackSys)
	case "Sys":
		return float64(met.Sys)
	case "TotalAlloc":
		return float64(met.TotalAlloc)
	default:
		return 0
	}
}
