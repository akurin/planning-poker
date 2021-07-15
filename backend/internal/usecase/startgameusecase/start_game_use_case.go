package startgameusecase

import "backend/internal/domain"

type UseCase interface {
	Execute() (domain.GameId, error)
}

func New(gameRepository domain.GameRepository) UseCase {
	return &startGameUseCase{
		gameRepository: gameRepository,
	}
}

type startGameUseCase struct {
	gameRepository domain.GameRepository
}

func (f *startGameUseCase) Execute() (domain.GameId, error) {
	game := domain.NewGame()

	err := f.gameRepository.Save(game)
	if err != nil {
		return nil, err
	}
	return game.Id(), nil
}
