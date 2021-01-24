package httpapi

import (
	"backend/internal/domain"
	"backend/internal/usecase/findgame"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GameApi struct {
	findGame findgame.FindGame
}

func NewGameApi(findGame findgame.FindGame) *GameApi {
	return &GameApi{findGame: findGame}
}

func (a *GameApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/games/{id}", a.getPlayer).Methods("GET")
}

func (a *GameApi) getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	game, err := a.findGame.Execute(domain.GameId(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if game == nil {
		http.NotFound(w, r)
		return
	}
	result := gameDto{
		Id: string(game.Id()),
	}
	_ = json.NewEncoder(w).Encode(result)
}
