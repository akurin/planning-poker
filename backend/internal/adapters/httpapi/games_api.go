package httpapi

import (
	"backend/internal/adapters/httpapi/httpconfig"
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

type GamesApi struct {
	config httpconfig.HttpApiConfig
}

func NewGamesApi(config httpconfig.HttpApiConfig) *GamesApi {
	return &GamesApi{
		config: config,
	}
}

func (a *GamesApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/games", a.post).Methods("POST")
}

func (a *GamesApi) post(w http.ResponseWriter, _ *http.Request) {
	location := a.config.BasePath()
	location.Path = path.Join(location.Path, "/games/1")
	w.Header().Set("Location", location.String())
	w.WriteHeader(http.StatusCreated)
}
