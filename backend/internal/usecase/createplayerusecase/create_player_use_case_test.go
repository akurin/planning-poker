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
	// arrange
	stubbedPlayerId := domain.NewFakePlayerId("some player id")
	playerIdGenerator := domain.NewFakePlayerIdGenerator(stubbedPlayerId)

	expectedPlayer := domain.NewPlayer(stubbedPlayerId, "John Doe")
	playerRepository := inmemoryrepo.NewPlayerRepository()

	sut := New(playerIdGenerator, playerRepository)

	// act
	createdPlayerId, err := sut.Execute("John Doe")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, stubbedPlayerId, createdPlayerId)

	createdPlayer, err := playerRepository.FindById(expectedPlayer.Id())
	require.NoError(t, err)
	assert.Equal(t, expectedPlayer, createdPlayer)
}

func Test_Create_player_with_broken_repository(t *testing.T) {
	// arrange
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()
	playerRepository := brokenrepo.NewPlayerRepository()
	sut := New(playerIdGenerator, playerRepository)

	// act
	createdUserId, err := sut.Execute("some")

	// assert
	assert.Empty(t, createdUserId)
	assert.NotNil(t, err)
}
