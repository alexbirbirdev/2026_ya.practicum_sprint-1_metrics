package server

import "github.com/caarlos0/env/v6"

type envConfig struct {
	Addr string `env:"ADDRESS,required"`
}

func applyEnv(cfg *Server) {
	var e envConfig
	err := env.Parse(&e)
	if err != nil {
		return
	}
	if e.Addr != "" {
		cfg.HTTP.Addr = e.Addr
	}
}
