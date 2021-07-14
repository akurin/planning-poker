package findgame

import (
	"backend/internal/adapters/brokenrepository"
	"backend/internal/adapters/inmemory"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Find_Game(t *testing.T) {
	gameId := domain.GameId("some")
	givenGame := domain.NewGame(gameId)
	gameRepository := inmemory.NewGameRepository()
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
	sut := NewFindGame(brokenrepository.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.NotNil(t, err)
	assert.Nil(t, game)
}

func Test_Find_Non_Existent_Game(t *testing.T) {
	gameId := domain.GameId("some")
	sut := NewFindGame(inmemory.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.Nil(t, err)
	assert.Nil(t, game)
}
