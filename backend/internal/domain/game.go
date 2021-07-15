package domain

import "github.com/google/uuid"

type GameId interface {
	String() string
}

func ParseGameId(s string) (GameId, error) {
	return uuid.Parse(s)
}

func NewGameId() GameId {
	return uuid.New()
}

type Game struct {
	id      GameId
	players []Player
}

func (g *Game) Id() GameId {
	return g.id
}

func (g *Game) addPlayer(p Player) {
	g.players = append(g.players, p)
}

func NewGame() *Game {
	return &Game{
		id:      NewGameId(),
		players: []Player{},
	}
}

func NewGameWithId(id GameId) *Game {
	return &Game{
		id:      id,
		players: []Player{},
	}
}
