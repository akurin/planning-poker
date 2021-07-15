package gameapi

import (
	"backend/internal/adapters/httpapi/dtos"
	"backend/internal/domain"
	"backend/internal/usecase/findgame"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GameApi struct {
	findGame findgame.UseCase
}

func NewGameApi(findGame findgame.UseCase) *GameApi {
	return &GameApi{findGame: findGame}
}

func (a *GameApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/games/{id}", a.getPlayer).Methods("GET")
}

func (a *GameApi) getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	parsed, err := domain.ParseGameId(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	game, err := a.findGame.Execute(parsed)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if game == nil {
		http.NotFound(w, r)
		return
	}
	result := dtos.GameDto{
		Id: game.Id().String(),
	}
	_ = json.NewEncoder(w).Encode(result)
}
