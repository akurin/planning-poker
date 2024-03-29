package joingameusecase

import (
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Join_game(t *testing.T) {
	// Arrange
	player := domain.NewPlayer(domain.NewFakePlayerId("some"), "name")
	playerRepository := inmemoryrepo.NewPlayerRepository()
	err := playerRepository.Save(player)
	require.NoError(t, err)

	gameRepository := inmemoryrepo.NewGameRepository()
	game := domain.NewGame()
	err = gameRepository.Save(game)
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// Act
	err = sut.Execute(player.Id(), game.Id())

	// Assert
	assert.NoError(t, err)
	j, err := gameRepository.FindById(game.Id())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(j.Players()))
}

func Test_Join_game_when_player_not_found(t *testing.T) {
	// Arrange
	playerRepository := inmemoryrepo.NewPlayerRepository()
	playerId := domain.NewFakePlayerId("some player id")

	gameRepository := inmemoryrepo.NewGameRepository()
	game := domain.NewGame()
	err := gameRepository.Save(game)
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// Act
	err = sut.Execute(playerId, game.Id())

	// Assert
	assert.Equal(t, NewPlayerNotFoundError(playerId), err)
}

func Test_Join_game_when_game_not_found(t *testing.T) {
	// Arrange
	player := domain.NewPlayer(domain.NewFakePlayerId("1"), "name")
	playerRepository := inmemoryrepo.NewPlayerRepository()
	err := playerRepository.Save(player)
	require.NoError(t, err)

	gameRepository := inmemoryrepo.NewGameRepository()
	gameId, err := domain.ParseGameId("b06d89ce-4be5-4f19-9e69-04e79a83c6c1")
	require.NoError(t, err)

	sut := New(playerRepository, gameRepository)

	// Act
	err = sut.Execute(player.Id(), gameId)

	// Assert
	assert.Equal(t, NewGameNotFoundError(gameId), err)
}
