package domain

type PlayerRepository interface {
	Save(player *Player) error
	FindById(playerId PlayerId) (*Player, error)
}
