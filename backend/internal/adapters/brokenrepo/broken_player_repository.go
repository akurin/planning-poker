package brokenrepo

import (
	"backend/internal/domain"
	"errors"
)

type brokenPlayerRepository struct {
}

func NewPlayerRepository() domain.PlayerRepository {
	return brokenPlayerRepository{}
}

func (r brokenPlayerRepository) Save(_ *domain.Player) error {
	return errors.New("unable to save")
}

func (r brokenPlayerRepository) FindById(_ domain.PlayerId) (*domain.Player, error) {
	return nil, errors.New("unable to find player by id")
}
