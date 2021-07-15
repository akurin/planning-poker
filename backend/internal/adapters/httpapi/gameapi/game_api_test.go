package gameapi

import (
	"backend/internal/domain"
	"backend/internal/usecase/findgameusecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get_a_game(t *testing.T) {
	findGameUseCase := findgameusecase.Mock()
	gameId, _ := domain.ParseGameId("b06d89ce-4be5-4f19-9e69-04e79a83c6c1")
	game := domain.NewGameWithId(gameId)
	findGameUseCase.ReturnGame(game)

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/b06d89ce-4be5-4f19-9e69-04e79a83c6c1", nil)
	rr := httptest.NewRecorder()
	wantBody := `{
		"id": "b06d89ce-4be5-4f19-9e69-04e79a83c6c1"
	}`

	handleWithGameApi(sut, rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, wantBody, rr.Body.String())
}

func Test_Game_not_found(t *testing.T) {
	findGameUseCase := findgameusecase.Mock()

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/b06d89ce-4be5-4f19-9e69-04e79a83c6c1", nil)
	rr := httptest.NewRecorder()

	handleWithGameApi(sut, rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "404 page not found\n", rr.Body.String())
}

func Test_Get_a_game_by_invalid_id(t *testing.T) {
	findGameUseCase := findgameusecase.Mock()

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/not-a-uuid", nil)
	rr := httptest.NewRecorder()

	handleWithGameApi(sut, rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, "404 page not found\n", rr.Body.String())
}

func Test_Get_a_game_when_use_case_fails(t *testing.T) {
	findGameUseCase := findgameusecase.Mock()
	findGameUseCase.ReturnError()

	sut := NewGameApi(findGameUseCase)
	req, _ := http.NewRequest("GET", "/games/b06d89ce-4be5-4f19-9e69-04e79a83c6c1", nil)
	rr := httptest.NewRecorder()

	handleWithGameApi(sut, rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error\n", rr.Body.String())
}

func handleWithGameApi(api *GameApi, w http.ResponseWriter, req *http.Request) {
	router := mux.NewRouter()
	api.AddRoutes(router)

	handler := http.HandlerFunc(router.ServeHTTP)
	handler.ServeHTTP(w, req)
}
