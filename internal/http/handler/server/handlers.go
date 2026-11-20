package server

import "net/http"

type ServerUsecase interface {
	Ping() error
}

type ServerHandlers struct {
	serverUsecases ServerUsecase
}

func NewServerHandler(
	serverUsecases ServerUsecase,
) *ServerHandlers {
	return &ServerHandlers{
		serverUsecases: serverUsecases,
	}
}

func (h ServerHandlers) Ping(res http.ResponseWriter, _ *http.Request) {
	err := h.serverUsecases.Ping()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
