package domain

type Player struct {
	id   PlayerId
	name string
}

func NewPlayer(id PlayerId, name string) *Player {
	return &Player{id: id, name: name}
}

func (p *Player) Id() PlayerId {
	return p.id
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Join(game *Game) {
	game.AddPlayer(p)
}

type PlayerId string
