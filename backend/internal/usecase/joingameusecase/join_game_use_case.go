package joingameusecase

import "backend/internal/domain"

type UseCase interface {
	Execute(playerId domain.PlayerId, gameId domain.GameId) error
}

func New(playerRepository domain.PlayerRepository, gameRepository domain.GameRepository) UseCase {
	return &joinGameUseCase{
		playerRepository: playerRepository,
		gameRepository:   gameRepository,
	}
}

type joinGameUseCase struct {
	playerRepository domain.PlayerRepository
	gameRepository   domain.GameRepository
}

func (j *joinGameUseCase) Execute(playerId domain.PlayerId, gameId domain.GameId) error {
	player, err := j.playerRepository.FindById(playerId)
	if err != nil {
		return err
	}
	if player == nil {
		return NewPlayerNotFoundError(playerId)
	}

	game, err := j.gameRepository.FindById(gameId)
	if err != nil {
		return err
	}
	if game == nil {
		return NewGameNotFoundError(gameId)
	}

	player.Join(game)

	return nil
}
