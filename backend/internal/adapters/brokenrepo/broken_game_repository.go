package brokenrepo

import (
	"backend/internal/domain"
	"errors"
)

type brokenGameRepository struct {
}

func NewGameRepository() domain.GameRepository {
	return brokenGameRepository{}
}

func (r brokenGameRepository) Save(_ *domain.Game) error {
	return errors.New("unable to save")
}

func (r brokenGameRepository) FindById(_ domain.GameId) (*domain.Game, error) {
	return nil, errors.New("unable to find player by id")
}
