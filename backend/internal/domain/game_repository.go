package domain

type GameRepository interface {
	Save(game *Game) error
	FindById(gameId GameId) (*Game, error)
}
