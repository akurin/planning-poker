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
	givenGame := domain.NewGame()
	gameRepository := inmemoryrepo.NewGameRepository()
	_ = gameRepository.Save(givenGame)
	sut := New(gameRepository)

	foundGame, err := sut.Execute(givenGame.Id())

	assert.Nil(t, err)
	if !reflect.DeepEqual(foundGame, givenGame) {
		t.Errorf("Got %v, want %v", foundGame, givenGame)
	}
}

func Test_Find_Game_Fails(t *testing.T) {
	gameId := domain.NewGameId()
	sut := New(brokenrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.NotNil(t, err)
	assert.Nil(t, game)
}

func Test_Find_Non_Existent_Game(t *testing.T) {
	gameId := domain.NewGameId()
	sut := New(inmemoryrepo.NewGameRepository())

	game, err := sut.Execute(gameId)

	assert.Nil(t, err)
	assert.Nil(t, game)
}
