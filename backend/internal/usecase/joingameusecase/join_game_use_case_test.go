package joingameusecase

import (
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Join_game(t *testing.T) {
	// arrange
	player := domain.NewPlayer("1", "name")
	playerRepository := inmemoryrepo.NewPlayerRepository()
	err := playerRepository.Save(player)
	require.NoError(t, err)

	gameRepository := inmemoryrepo.NewGameRepository()
	game := domain.NewGame()
	err = gameRepository.Save(game)
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// act
	err = sut.Execute(player.Id(), game.Id())

	// assert
	assert.NoError(t, err)
	j, err := gameRepository.FindById(game.Id())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(j.Players()))
}

func Test_Join_game_when_player_not_found(t *testing.T) {
	// arrange
	playerRepository := inmemoryrepo.NewPlayerRepository()
	playerId := domain.PlayerId("some player id")

	gameRepository := inmemoryrepo.NewGameRepository()
	game := domain.NewGame()
	err := gameRepository.Save(game)
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// act
	err = sut.Execute(playerId, game.Id())

	// assert
	assert.Equal(t, NewPlayerNotFoundError(playerId), err)
}

func Test_Join_game_when_game_not_found(t *testing.T) {
	// arrange
	player := domain.NewPlayer("1", "name")
	playerRepository := inmemoryrepo.NewPlayerRepository()
	err := playerRepository.Save(player)
	require.NoError(t, err)

	gameRepository := inmemoryrepo.NewGameRepository()
	gameId, err := domain.ParseGameId("b06d89ce-4be5-4f19-9e69-04e79a83c6c1")
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// act
	err = sut.Execute(player.Id(), gameId)

	// assert
	assert.Equal(t, NewGameNotFoundError(gameId), err)
}
