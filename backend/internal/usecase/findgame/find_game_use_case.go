package findgame

import "backend/internal/domain"

type UseCase interface {
	Execute(id domain.GameId) (*domain.Game, error)
}

func New(gameRepository domain.GameRepository) UseCase {
	return &findGameUseCase{
		gameRepository: gameRepository,
	}
}

type findGameUseCase struct {
	gameRepository domain.GameRepository
}

func (f *findGameUseCase) Execute(id domain.GameId) (*domain.Game, error) {
	game, err := f.gameRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return game, nil
}
