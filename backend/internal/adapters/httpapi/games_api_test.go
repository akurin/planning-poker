package httpapi

import (
	"backend/internal/adapters/httpapi/httpconfig"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_Start_a_new_game(t *testing.T) {
	basePath, _ := url.Parse("http://mydomain")
	config := httpconfig.NewHttpApiConfig(basePath)
	sut := NewGamesApi(config)

	reqBody := `{ "title": "Sprint 23 planning" }`
	req, _ := http.NewRequest("POST", "/games", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	handleWithGamesApi(sut, rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	location := rr.Header().Get("location")
	assert.Equal(t, "http://mydomain/games/1", location)
}

func handleWithGamesApi(api *GamesApi, w http.ResponseWriter, req *http.Request) {
	router := mux.NewRouter()
	api.AddRoutes(router)

	handler := http.HandlerFunc(router.ServeHTTP)
	handler.ServeHTTP(w, req)
}
