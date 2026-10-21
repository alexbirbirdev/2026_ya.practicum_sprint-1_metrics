package server

import "flag"

func parseByFlag() Server {
	addr := flag.String("a", "localhost:8080", "адрес эндопинта HTTP сервера")
	flag.Parse()

	return Server{
		HTTP: HTTPServer{Addr: *addr},
	}
}
