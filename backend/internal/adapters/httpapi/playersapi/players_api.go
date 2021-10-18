package playersapi

import (
	"backend/internal/adapters/httpapi/dtos"
	"backend/internal/usecase/createplayerusecase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type PlayersApi struct {
	basePath            *url.URL
	createPlayerUseCase createplayerusecase.UseCase
}

func New(basePath *url.URL, createPlayerUseCase createplayerusecase.UseCase) *PlayersApi {
	return &PlayersApi{
		basePath:            basePath,
		createPlayerUseCase: createPlayerUseCase,
	}
}

func (a *PlayersApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/players", a.post).Methods("POST")
}

func (a *PlayersApi) post(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var createPlayerDto dtos.CreatePlayerDto
	err := json.Unmarshal(reqBody, &createPlayerDto)
	if err != nil {
		// todo
	}

	if validationResult := createPlayerDto.Validate(); validationResult != "" {
		http.Error(w, validationResult, http.StatusBadRequest)
		return
	}

	result, err := a.createPlayerUseCase.Execute(createPlayerDto.Name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	location := *a.basePath
	location.Path = path.Join(location.Path, fmt.Sprintf("/players/%s", result.Id()))
	w.Header().Set("Location", location.String())

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.Token(),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}
