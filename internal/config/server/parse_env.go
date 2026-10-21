package server

import "github.com/caarlos0/env/v6"

type envConfig struct {
	Addr string `env:"ADDRESS,required"`
}

func parseByEnv() Server {
	var cfg envConfig
	err := env.Parse(&cfg)
	if err != nil {
		return Server{}
	}
	return Server{
		HTTP: HTTPServer{cfg.Addr},
	}
}
