package signupapi

import (
	"backend/internal/adapters/httpapi/dtos"
	"backend/internal/adapters/httpapi/sessions"
	"backend/internal/usecase/createplayerusecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type SignupApi struct {
	createPlayerUseCase createplayerusecase.UseCase
	sessionStore        sessions.SessionStore
}

func New(createPlayerUseCase createplayerusecase.UseCase, sessionStore sessions.SessionStore) *SignupApi {
	return &SignupApi{
		createPlayerUseCase: createPlayerUseCase,
		sessionStore:        sessionStore,
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
		w.WriteHeader(http.StatusBadRequest)
		return
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

	session, err := a.sessionStore.Get(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session.Authenticate(result.Id())

	err = a.sessionStore.Save(w, r, session)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
