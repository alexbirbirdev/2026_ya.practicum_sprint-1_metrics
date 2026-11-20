package server

import (
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
)

type envConfig struct {
	Addr            string `env:"ADDRESS"`
	StoreInterval   int    `env:"STORE_INTERVAL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	Restore         bool   `env:"RESTORE"`
	DSN             string `env:"DATABASE_DSN"`
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
	if _, ok := os.LookupEnv("STORE_INTERVAL"); ok {
		if e.StoreInterval >= 0 {
			cfg.FileStorage.StoreInterval = time.Duration(e.StoreInterval) * time.Second
		}
	}
	if _, ok := os.LookupEnv("FILE_STORAGE_PATH"); ok && e.FileStoragePath != "" {
		cfg.FileStorage.FileStoragePath = e.FileStoragePath
	}
	if _, ok := os.LookupEnv("RESTORE"); ok {
		cfg.FileStorage.Restore = parseBool(os.Getenv("RESTORE"))
	}
	if _, ok := os.LookupEnv("DATABASE_DSN"); ok && e.DSN != "" {
		cfg.DB.DSN = e.DSN
	}
}

func parseBool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true":
		return true
	default:
		return false
	}
}
