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

type Server struct {
	HTTP        HTTPServer
	FileStorage FileStorage
}

func GetConfig() Server {
	cfg := parseByFlag()
	applyEnv(&cfg)
	return cfg
}
