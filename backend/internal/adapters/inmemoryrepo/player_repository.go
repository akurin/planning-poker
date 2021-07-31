package inmemoryrepo

import "backend/internal/domain"

type playerRepository struct {
	players map[domain.PlayerId]*domain.Player
}

func NewPlayerRepository() domain.PlayerRepository {
	return &playerRepository{
		players: map[domain.PlayerId]*domain.Player{},
	}
}

func (r *playerRepository) Save(player *domain.Player) error {
	r.players[player.Id()] = player
	return nil
}

func (r *playerRepository) FindById(playerId domain.PlayerId) (*domain.Player, error) {
	player, ok := r.players[playerId]
	if ok {
		return player, nil
	} else {
		return nil, nil
	}
}
