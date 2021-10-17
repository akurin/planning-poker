package createplayerusecase

import "backend/internal/domain"

type UseCase interface {
	Execute(name string) (Result, error)
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

func (c *createPlayerUseCase) Execute(name string) (Result, error) {
	createdPlayerId := c.playerIdGenerator.Generate()
	createdPlayer := domain.NewPlayer(createdPlayerId, name)

	err := c.playerRepository.Save(createdPlayer)
	if err != nil {
		return Result{}, err
	}
	return NewResult(createdPlayerId, "todo"), nil
}

// todo
type Result struct {
	id    domain.PlayerId
	token string
}

func NewResult(id domain.PlayerId, token string) Result {
	return Result{id, token}
}
