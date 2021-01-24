package httpapi

import (
	"backend/internal/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get_Game_When_Repository_Is_Broken(t *testing.T) {
	findGameUseCase := NewFindGameMock()
	findGameUseCase.ReturnError()

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/1", nil)
	rr := httptest.NewRecorder()

	handleWithApi(sut, rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error\n", rr.Body.String())
}

func handleWithApi(api *GameApi, w http.ResponseWriter, req *http.Request) {
	router := mux.NewRouter()
	api.AddRoutes(router)

	handler := http.HandlerFunc(router.ServeHTTP)
	handler.ServeHTTP(w, req)
}

func Test_Get_Non_Existent_Game(t *testing.T) {
	findGameUseCase := NewFindGameMock()

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/1", nil)
	rr := httptest.NewRecorder()

	handleWithApi(sut, rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "404 page not found\n", rr.Body.String())
}

func Test_Get_Existent_Game(t *testing.T) {
	findGameUseCase := NewFindGameMock()
	findGameUseCase.ReturnGame(domain.NewGame("1"))

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/1", nil)
	rr := httptest.NewRecorder()
	wantBody := `{
		"id": "1"
	}`

	handleWithApi(sut, rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	s := rr.Body.String()
	assert.JSONEq(t, wantBody, s)
}
