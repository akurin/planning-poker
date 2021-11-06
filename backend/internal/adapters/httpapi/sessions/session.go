package sessions

import "backend/internal/domain"

type Session interface {
	IsAuthenticated() bool
	PlayerId() domain.PlayerId
	Authenticate(playerId domain.PlayerId)
}
