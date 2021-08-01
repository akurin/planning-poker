package createplayerusecase

import "backend/internal/domain"

type UseCase interface {
	Execute(name string) (domain.PlayerId, error)
}

func New(playerIdGenerator domain.PlayerIdGenerator, playerRepository domain.PlayerRepository) UseCase {
	return &createPlayerUseCase{
		playerIdGenerator: playerIdGenerator,
		playerRepository:  playerRepository,
	}
}

type createPlayerUseCase struct {
	playerRepository  domain.PlayerRepository
	playerIdGenerator domain.PlayerIdGenerator
}

func (c *createPlayerUseCase) Execute(name string) (domain.PlayerId, error) {
	createdPlayerId := c.playerIdGenerator.Generate()
	createdPlayer := domain.NewPlayer(createdPlayerId, name)

	err := c.playerRepository.Save(createdPlayer)
	if err != nil {
		return nil, err
	}
	return createdPlayerId, nil
}
