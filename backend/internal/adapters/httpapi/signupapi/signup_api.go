package signupapi

import (
	"backend/internal/adapters/httpapi/dtos"
	"backend/internal/usecase/createplayerusecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type SignupApi struct {
	createPlayerUseCase createplayerusecase.UseCase
}

func New(createPlayerUseCase createplayerusecase.UseCase) *SignupApi {
	return &SignupApi{
		createPlayerUseCase: createPlayerUseCase,
	}
}

func (a *SignupApi) AddRoutes(router *mux.Router) {
	router.HandleFunc("/signup", a.post).Methods("POST")
}

func (a *SignupApi) post(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.Token(),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}
