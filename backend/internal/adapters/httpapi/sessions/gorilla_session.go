package sessions

import (
	"backend/internal/domain"
	"github.com/gorilla/sessions"
)

type gorillaSession struct {
	session *sessions.Session
}

func NewGorillaSession(session *sessions.Session) Session {
	return &gorillaSession{
		session: session,
	}
}

func (g *gorillaSession) IsAuthenticated() bool {
	ok, value := g.session.Values["authenticated"].(bool)
	return ok && value
}

func (g *gorillaSession) PlayerId() domain.PlayerId {
	value, ok := g.session.Values["player_id"].(string)
	if !ok {
		panic("no player_id in session")
	}
	playerId, err := domain.ParsePlayerId(value)
	if err != nil {
		panic("invalid player_id in session")
	}
	return playerId
}

func (g *gorillaSession) Authenticate(playerId domain.PlayerId) {
	g.session.Values["authenticated"] = true
	g.session.Values["player_id"] = playerId.String()
}
