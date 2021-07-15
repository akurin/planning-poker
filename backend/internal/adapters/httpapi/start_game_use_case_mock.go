package httpapi

import (
	"backend/internal/domain"
	"errors"
)

type startGameUseCaseMock struct {
	gameId domain.GameId
	error  error
}

func NewStartGameUseCaseMock() *startGameUseCaseMock {
	return &startGameUseCaseMock{}
}

func (m *startGameUseCaseMock) Execute() (domain.GameId, error) {
	return m.gameId, m.error
}

func (m *startGameUseCaseMock) ReturnError() {
	m.gameId = nil
	m.error = errors.New("some error")
}

func (m *startGameUseCaseMock) ReturnGameId(gameId domain.GameId) {
	m.gameId = gameId
	m.error = nil
}
