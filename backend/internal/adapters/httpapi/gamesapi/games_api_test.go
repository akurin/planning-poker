package gamesapi

import (
	"backend/internal/domain"
	"backend/internal/usecase/startgameusecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_Start_a_new_game(t *testing.T) {
	basePath, err := url.Parse("https://mydomain")
	require.NoError(t, err)

	gameId, err := domain.ParseGameId("b06d89ce-4be5-4f19-9e69-04e79a83c6c1")
	require.NoError(t, err)

	startGameUseCaseMock := startgameusecase.Mock()
	startGameUseCaseMock.ReturnGameId(gameId)

	sut := NewGamesApi(basePath, startGameUseCaseMock)

	reqBody := `{ "title": "Sprint 23 planning" }`
	req, err := http.NewRequest("POST", "/games", strings.NewReader(reqBody))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	handleWithGamesApi(sut, rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	location := rr.Header().Get("Location")
	assert.Equal(t, "https://mydomain/games/b06d89ce-4be5-4f19-9e69-04e79a83c6c1", location)
}

func Test_Start_a_new_game_when_use_case_fails(t *testing.T) {
	basePath, err := url.Parse("https://mydomain")
	require.NoError(t, err)

	startGameUseCaseMock := startgameusecase.Mock()
	startGameUseCaseMock.ReturnError()

	sut := NewGamesApi(basePath, startGameUseCaseMock)

	reqBody := `{ "title": "Sprint 23 planning" }`
	req, err := http.NewRequest("POST", "/games", strings.NewReader(reqBody))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	handleWithGamesApi(sut, rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error\n", rr.Body.String())
}

func handleWithGamesApi(api *GamesApi, w http.ResponseWriter, req *http.Request) {
	router := mux.NewRouter()
	api.AddRoutes(router)

	handler := http.HandlerFunc(router.ServeHTTP)
	handler.ServeHTTP(w, req)
}
