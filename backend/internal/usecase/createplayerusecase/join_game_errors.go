package createplayerusecase

import (
	"backend/internal/domain"
	"fmt"
)

type PlayerNotFoundError struct {
	playerId domain.PlayerId
}

func NewPlayerNotFoundError(playerId domain.PlayerId) PlayerNotFoundError {
	return PlayerNotFoundError{
		playerId: playerId,
	}
}

func (e PlayerNotFoundError) Error() string {
	return fmt.Sprintf("Unable to find player %s", e.playerId)
}

type GameNotFoundError struct {
	gameId domain.GameId
}

func NewGameNotFoundError(gameId domain.GameId) GameNotFoundError {
	return GameNotFoundError{
		gameId: gameId,
	}
}

func (e GameNotFoundError) Error() string {
	return fmt.Sprintf("Unable to find game %s", e.gameId)
}
