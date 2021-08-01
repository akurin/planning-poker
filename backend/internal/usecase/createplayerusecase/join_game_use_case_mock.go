package createplayerusecase

import (
	"backend/internal/domain"
	"errors"
)

type createPlayerUseCaseMock struct {
	playerId domain.PlayerId
	error    error
}

func Mock() *createPlayerUseCaseMock {
	return &createPlayerUseCaseMock{
		playerId: domain.NewFakePlayerId("some"),
	}
}

func (m *createPlayerUseCaseMock) Execute(_ string) (domain.PlayerId, error) {
	return m.playerId, m.error
}

func (m *createPlayerUseCaseMock) ReturnError() {
	m.playerId = nil
	m.error = errors.New("some error")
}

func (m *createPlayerUseCaseMock) ReturnPlayerId(gameId domain.PlayerId) {
	m.playerId = gameId
	m.error = nil
}
