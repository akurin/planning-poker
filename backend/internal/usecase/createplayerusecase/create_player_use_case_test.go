package createplayerusecase

import (
	"backend/internal/adapters/brokenrepo"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Create_player(t *testing.T) {
	// Arrange
	stubbedPlayerId := domain.NewFakePlayerId("some player id")
	playerIdGenerator := domain.NewFakePlayerIdGenerator(stubbedPlayerId)

	expectedPlayer := domain.NewPlayer(stubbedPlayerId, "John Doe")
	playerRepository := inmemoryrepo.NewPlayerRepository()

	sut := New(playerIdGenerator, playerRepository)

	// Act
	result, err := sut.Execute("John Doe")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, stubbedPlayerId, result.id)

	createdPlayer, err := playerRepository.FindById(expectedPlayer.Id())
	require.NoError(t, err)
	assert.Equal(t, expectedPlayer, createdPlayer)
}

func Test_Create_player_with_broken_repository(t *testing.T) {
	// Arrange
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()
	playerRepository := brokenrepo.NewPlayerRepository()
	sut := New(playerIdGenerator, playerRepository)

	// Act
	createdUserId, err := sut.Execute("some")

	// Assert
	assert.Empty(t, createdUserId)
	assert.NotNil(t, err)
}
