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

func parseByEnv() Agent {
	var cfg envConfig
	err := env.Parse(&cfg)
	if err != nil {
		return Agent{}
	}
	return Agent{
		HTTP: HTTPClient{
			Addr: cfg.Addr,
		},
		Collector: Collector{
			PollInterval: time.Duration(cfg.PollInterval) * time.Second,
		},
		Reporter: Reporter{
			ReportInterval: time.Duration(cfg.ReportInterval) * time.Second,
		},
	}
}
