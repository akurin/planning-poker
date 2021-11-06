package sessionauth

import (
	"backend/internal/adapters/httpapi/sessionauth/fakehandler"
	"backend/internal/adapters/httpapi/sessions"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Authenticated_session(t *testing.T) {
	// Arrange
	nextHandler := fakehandler.New()
	sessionStore := sessions.NewFakeStore()
	sut := NewSessionMiddleware(nextHandler, sessionStore)

	rr := httptest.NewRecorder()
	req := someRequest(t)

	session, err := sessionStore.Get(req)
	assert.NoError(t, err)
	session.Authenticate(domain.NewFakePlayerId("some-player-id"))

	// Act
	sut.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, 200, rr.Code)
	assert.True(t, nextHandler.WasCalled())
}

func someRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest("GET", "/some", nil)
	assert.NoError(t, err)
	return req
}

func Test_Missing_session(t *testing.T) {
	// Arrange
	nextHandler := fakehandler.New()
	sessionStore := sessions.NewFakeStore()
	sut := NewSessionMiddleware(nextHandler, sessionStore)
	rr := httptest.NewRecorder()
	req := someRequest(t)

	// Act
	sut.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, 401, rr.Code)
	assert.False(t, nextHandler.WasCalled())
}
