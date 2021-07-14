package findgame

import (
	"backend/internal/adapters/brokenrepo"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Find_Game(t *testing.T) {
	gameId := domain.GameId("some")
	givenGame := domain.NewGame(gameId)
	gameRepository := inmemoryrepo.NewGameRepository()
	_ = gameRepository.Save(givenGame)
	sut := NewFindGame(gameRepository)

	foundGame, err := sut.Execute(gameId)

	assert.Nil(t, err)
	if !reflect.DeepEqual(foundGame, givenGame) {
		t.Errorf("Got %v, want %v", foundGame, givenGame)
	}
}

func Test_Find_Game_Fails(t *testing.T) {
	gameId := domain.GameId("some")
	sut := NewFindGame(brokenrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.NotNil(t, err)
	assert.Nil(t, game)
}

func Test_Find_Non_Existent_Game(t *testing.T) {
	gameId := domain.GameId("some")
	sut := NewFindGame(inmemoryrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.Nil(t, err)
	assert.Nil(t, game)
}
