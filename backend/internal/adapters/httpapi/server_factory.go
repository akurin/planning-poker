package httpapi

import (
	"backend/internal/adapters/httpapi/gameapi"
	"github.com/gorilla/mux"
	"net/http"
)

type ServerFactory struct {
	playerApi *gameapi.GameApi
}

func NewServerFactory(playerApi *gameapi.GameApi) *ServerFactory {
	return &ServerFactory{playerApi: playerApi}
}

func (s *ServerFactory) NewServer() *http.Server {
	router := mux.NewRouter()
	s.playerApi.AddRoutes(router)

	handler := http.NewServeMux()
	handler.Handle("/", router)

	// TODO: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://habr.com/ru/post/197468/
	return &http.Server{Addr: ":8181", Handler: handler}
}
