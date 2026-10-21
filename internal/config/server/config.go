package server

import "flag"

type HTTPServer struct {
	Addr string
}

type Server struct {
	HTTP HTTPServer
}

func ParseServer() Server {
	addr := flag.String("a", "localhost:8080", "адрес эндопинта HTTP сервера")
	flag.Parse()

	return Server{
		HTTP: HTTPServer{Addr: *addr},
	}
}
