package server

import "time"

type HTTPServer struct {
	Addr string
}

type FileStorage struct {
	StoreInterval   time.Duration
	FileStoragePath string
	Restore         bool
}
type DB struct {
	DSN string
}

type Server struct {
	HTTP        HTTPServer
	FileStorage FileStorage
	DB          DB
}

func GetConfig() Server {
	cfg := parseByFlag()
	applyEnv(&cfg)
	return cfg
}
