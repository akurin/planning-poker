package inmemory

import "backend/internal/domain"

type gameRepository struct {
	games map[domain.GameId]*domain.Game
}

func NewGameRepository() domain.GameRepository {
	return &gameRepository{
		games: map[domain.GameId]*domain.Game{},
	}
}

func (r *gameRepository) Save(game *domain.Game) error {
	r.games[game.Id()] = game
	return nil
}

func (r *gameRepository) FindById(gameId domain.GameId) (*domain.Game, error) {
	game, ok := r.games[gameId]
	if ok {
		return game, nil
	} else {
		return nil, nil
	}
}
