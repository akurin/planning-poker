package httpapi

import (
	"backend/internal/adapters/httpapi/gameapi"
	"backend/internal/adapters/httpapi/gamesapi"
	"backend/internal/adapters/httpapi/signupapi"
	"github.com/gorilla/mux"
	"net/http"
)

type ServerFactory struct {
	playersApi *signupapi.SignupApi
	gamesApi   *gamesapi.GamesApi
	gameApi    *gameapi.GameApi
}

func NewServerFactory(
	playersApi *signupapi.SignupApi,
	gamesApi *gamesapi.GamesApi,
	gameApi *gameapi.GameApi,
) *ServerFactory {
	return &ServerFactory{
		playersApi: playersApi,
		gamesApi:   gamesApi,
		gameApi:    gameApi,
	}
}

func (s *ServerFactory) NewServer() *http.Server {
	router := mux.NewRouter()
	s.playersApi.AddRoutes(router)
	s.gamesApi.AddRoutes(router)
	s.gameApi.AddRoutes(router)

	handler := http.NewServeMux()
	handler.Handle("/", router)

	// TODO: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://habr.com/ru/post/197468/
	return &http.Server{Addr: ":8181", Handler: handler}
}
