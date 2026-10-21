package server

type HTTPServer struct {
	Addr string
}

type Server struct {
	HTTP HTTPServer
}

func GetConfig() Server {
	var emptyConfig Server
	cfg := parseByEnv()

	if cfg != emptyConfig {
		return cfg
	}

	cfg = parseByFlag()
	return cfg
}
