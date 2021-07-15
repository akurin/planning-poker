package gamesapi

import (
	"backend/internal/adapters/httpapi"
	"backend/internal/usecase/startgame"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

type GamesApi struct {
	config           httpapi.HttpApiConfig
	startGameUseCase startgame.UseCase
}

func NewGamesApi(config httpapi.HttpApiConfig, startGameUseCase startgame.UseCase) *GamesApi {
	return &GamesApi{
		config:           config,
		startGameUseCase: startGameUseCase,
	}
}

func (a *GamesApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/games", a.post).Methods("POST")
}

func (a *GamesApi) post(w http.ResponseWriter, _ *http.Request) {
	gameId, err := a.startGameUseCase.Execute()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	location := a.config.BasePath()
	location.Path = path.Join(location.Path, fmt.Sprintf("/games/%s", gameId))
	w.Header().Set("Location", location.String())
	w.WriteHeader(http.StatusCreated)
}
