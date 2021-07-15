package findgameusecase

import (
	"backend/internal/domain"
	"errors"
)

type findGameUseCaseMock struct {
	game  *domain.Game
	error error
}

func Mock() *findGameUseCaseMock {
	return &findGameUseCaseMock{}
}

func (m *findGameUseCaseMock) Execute(_ domain.GameId) (*domain.Game, error) {
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
