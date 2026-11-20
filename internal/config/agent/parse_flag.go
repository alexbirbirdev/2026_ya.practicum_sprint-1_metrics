package agent

import (
	"flag"
	"time"
)

func parseByFlag() Agent {
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
