package agent

import (
	"time"
)

type HTTPClient struct {
	Addr string
}

func (c HTTPClient) BaseURL() string {
	return "http://" + c.Addr
}

type Collector struct {
	PollInterval time.Duration
}

type Reporter struct {
	ReportInterval time.Duration
}

type Agent struct {
	HTTP      HTTPClient
	Collector Collector
	Reporter  Reporter
}

func GetConfig() Agent {
	var emptyConfig Agent
	cfg := parseByEnv()

	if cfg != emptyConfig {
		return cfg
	}

	cfg = parseByFlag()
	return cfg
}
