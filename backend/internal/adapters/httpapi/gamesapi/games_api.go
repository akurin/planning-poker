package gamesapi

import (
	"backend/internal/usecase/startgameusecase"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"path"
)

type GamesApi struct {
	basePath         *url.URL
	startGameUseCase startgameusecase.UseCase
}

func NewGamesApi(basePath *url.URL, startGameUseCase startgameusecase.UseCase) *GamesApi {
	return &GamesApi{
		basePath:         basePath,
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

	location := a.basePath
	location.Path = path.Join(location.Path, fmt.Sprintf("/games/%s", gameId))
	w.Header().Set("Location", location.String())
	w.WriteHeader(http.StatusCreated)
}
