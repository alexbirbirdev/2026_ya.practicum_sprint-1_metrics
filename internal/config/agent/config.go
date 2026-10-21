package agent

import (
	"flag"
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

func ParseAgent() Agent {
	addr := flag.String("a", "localhost:8080", "адрес эндопинта HTTP сервера")
	reportSec := flag.Int("r", 10, "частота отправки метрик на сервер")
	pollSec := flag.Int("p", 2, "частота опроса метрик")
	flag.Parse()

	return Agent{
		HTTP: HTTPClient{
			Addr: *addr,
		},
		Collector: Collector{
			PollInterval: time.Duration(*pollSec) * time.Second,
		},
		Reporter: Reporter{
			ReportInterval: time.Duration(*reportSec) * time.Second,
		},
	}
}
