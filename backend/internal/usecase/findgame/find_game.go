package findgame

import "backend/internal/domain"

type FindGame interface {
	Execute(id domain.GameId) (*domain.Game, error)
}

func NewFindGame() FindGame {
	return findGame{}
}

type findGame struct {
}

func (f findGame) Execute(id domain.GameId) (*domain.Game, error) {
	panic("implement me")
}
