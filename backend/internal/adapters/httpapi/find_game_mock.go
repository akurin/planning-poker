package httpapi

import (
	"backend/internal/domain"
	"errors"
)

type findGameMock struct {
	game  *domain.Game
	error error
}

func NewFindGameMock() *findGameMock {
	return &findGameMock{}
}

func (m *findGameMock) Execute(id domain.GameId) (*domain.Game, error) {
	return m.game, m.error
}

func (m *findGameMock) ReturnError() {
	m.game = nil
	m.error = errors.New("some error")
}

func (m *findGameMock) ReturnGame(game *domain.Game) {
	m.game = game
	m.error = nil
}
