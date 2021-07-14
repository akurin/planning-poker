package domain

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

type GameId string

func NewGame(id GameId) *Game {
	return &Game{
		id:      id,
		players: []Player{},
	}
}
