package findgame

import "backend/internal/domain"

type FindGame interface {
	Execute(id domain.GameId) (*domain.Game, error)
}

func NewFindGame(gameRepository domain.GameRepository) FindGame {
	return &findGame{
		gameRepository: gameRepository,
	}
}

type findGame struct {
	gameRepository domain.GameRepository
}

func (f *findGame) Execute(id domain.GameId) (*domain.Game, error) {
	game, err := f.gameRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}
