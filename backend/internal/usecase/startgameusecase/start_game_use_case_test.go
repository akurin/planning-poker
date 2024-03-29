package startgameusecase

import (
	"backend/internal/adapters/brokenrepo"
	"backend/internal/adapters/inmemoryrepo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Start_game(t *testing.T) {
	gameRepository := inmemoryrepo.NewGameRepository()
	sut := New(gameRepository)

	startedGameId, err := sut.Execute()

	startedGame, _ := gameRepository.FindById(startedGameId)
	assert.Nil(t, err)
	assert.NotNil(t, startedGame)
}

func Test_Broken_repository(t *testing.T) {
	gameRepository := brokenrepo.NewGameRepository()
	sut := New(gameRepository)

	startedGameId, err := sut.Execute()

	assert.NotNil(t, err)
	assert.Nil(t, startedGameId)
}
