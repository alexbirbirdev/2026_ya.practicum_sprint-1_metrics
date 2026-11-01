package agent

import (
	"time"

	"github.com/caarlos0/env/v6"
)

type envConfig struct {
	Addr           string `env:"ADDRESS,required"`
	PollInterval   int    `env:"POLL_INTERVAL,required"`
	ReportInterval int    `env:"REPORT_INTERVAL,required"`
}

func applyEnv(cfg *Agent) {
	var e envConfig
	err := env.Parse(&e)
	if err != nil {
		return
	}

	if e.Addr != "" {
		cfg.HTTP.Addr = e.Addr
	}
	if e.PollInterval > 0 {
		cfg.Collector.PollInterval = time.Duration(e.PollInterval) * time.Second
	}
	if e.ReportInterval > 0 {
		cfg.Reporter.ReportInterval = time.Duration(e.ReportInterval) * time.Second
	}
}
