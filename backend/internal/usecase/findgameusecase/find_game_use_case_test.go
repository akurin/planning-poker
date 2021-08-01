package findgameusecase

import (
	"backend/internal/adapters/brokenrepo"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Game_found(t *testing.T) {
	givenGame := domain.NewGame()
	gameRepository := inmemoryrepo.NewGameRepository()
	_ = gameRepository.Save(givenGame)
	sut := New(gameRepository)

	foundGame, err := sut.Execute(givenGame.Id())

	assert.Nil(t, err)
	assert.Equal(t, givenGame, foundGame)
}

func Test_Game_not_found(t *testing.T) {
	gameId := domain.NewGameId()
	sut := New(inmemoryrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.Nil(t, err)
	assert.Nil(t, game)
}

func Test_Broken_repository(t *testing.T) {
	gameId := domain.NewGameId()
	sut := New(brokenrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.NotNil(t, err)
	assert.Nil(t, game)
}
