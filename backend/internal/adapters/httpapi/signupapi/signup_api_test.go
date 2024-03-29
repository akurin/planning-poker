package signupapi

import (
	"backend/internal/adapters/httpapi/sessions"
	"backend/internal/domain"
	"backend/internal/usecase/createplayerusecase"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Sign_up(t *testing.T) {
	// Arrange
	playerId := domain.NewFakePlayerId("some-player-id")

	createPlayerUseCase := createplayerusecase.Mock(
		createplayerusecase.WithResult(
			createplayerusecase.NewResult(playerId)))

	sessionStore := sessions.NewFakeStore()
	sut := New(createPlayerUseCase, sessionStore)

	reqBody := `{ "name": "John Doe" }`
	req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	// Act
	handleWithGamesApi(sut, rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Regexp(t, "^session_id=", rr.Header().Get("Set-Cookie"))
}

func Test_Create_player_when_use_case_fails(t *testing.T) {
	// Arrange
	createPlayerUseCase := createplayerusecase.Mock(
		createplayerusecase.WithError(errors.New("some")))

	sessionStore := sessions.NewFakeStore()
	sut := New(createPlayerUseCase, sessionStore)

	reqBody := `{ "name": "John Doe" }`
	req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	// Act
	handleWithGamesApi(sut, rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Internal Server Error\n", rr.Body.String())
}

func Test_Create_player_with_invalid_request_body(t *testing.T) {
	// Arrange
	createPlayerUseCase := createplayerusecase.Mock()

	sessionStore := sessions.NewFakeStore()
	sut := New(createPlayerUseCase, sessionStore)

	reqBody := `{ "name": null }`
	req, err := http.NewRequest("POST", "/signup", strings.NewReader(reqBody))
	require.NoError(t, err)
	rr := httptest.NewRecorder()

	// Act
	handleWithGamesApi(sut, rr, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "The field name is required.\n", rr.Body.String())
}

func handleWithGamesApi(api *SignupApi, w http.ResponseWriter, req *http.Request) {
	router := mux.NewRouter()
	api.AddRoutes(router)

	handler := http.HandlerFunc(router.ServeHTTP)
	handler.ServeHTTP(w, req)
}
