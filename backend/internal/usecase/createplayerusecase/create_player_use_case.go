package createplayerusecase

import "backend/internal/domain"

type UseCase interface {
	Execute(name string) (Result, error)
}

func New(
	playerIdGenerator domain.PlayerIdGenerator,
	playerRepository domain.PlayerRepository,
	tokenService domain.TokenService,
) UseCase {
	return &createPlayerUseCase{
		playerIdGenerator: playerIdGenerator,
		playerRepository:  playerRepository,
		tokenService:      tokenService,
	}
}

type createPlayerUseCase struct {
	playerRepository  domain.PlayerRepository
	playerIdGenerator domain.PlayerIdGenerator
	tokenService      domain.TokenService
}

func (c *createPlayerUseCase) Execute(name string) (Result, error) {
	createdPlayerId := c.playerIdGenerator.Generate()
	createdPlayer := domain.NewPlayer(createdPlayerId, name)

	err := c.playerRepository.Save(createdPlayer)
	if err != nil {
		return Result{}, err
	}
	token := c.tokenService.IssueToken(createdPlayerId)
	return NewResult(createdPlayerId, token), nil
}

type Result struct {
	id    domain.PlayerId
	token string
}

func NewResult(id domain.PlayerId, token string) Result {
	return Result{id, token}
}
