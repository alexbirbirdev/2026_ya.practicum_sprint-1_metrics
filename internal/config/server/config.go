package server

type HTTPServer struct {
	Addr string
}

type Server struct {
	HTTP HTTPServer
}

func GetConfig() Server {
	cfg := parseByFlag()
	applyEnv(&cfg)
	return cfg
}
