package findgame

import (
	"backend/internal/domain"
	"errors"
)

type findGameUseCaseMock struct {
	game  *domain.Game
	error error
}

func NewFindGameUseCaseMock() *findGameUseCaseMock {
	return &findGameUseCaseMock{}
}

func (m *findGameUseCaseMock) Execute(id domain.GameId) (*domain.Game, error) {
	return m.game, m.error
}

func (m *findGameUseCaseMock) ReturnError() {
	m.game = nil
	m.error = errors.New("some error")
}

func (m *findGameUseCaseMock) ReturnGame(game *domain.Game) {
	m.game = game
	m.error = nil
}
