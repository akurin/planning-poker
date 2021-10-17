package createplayerusecase

import (
	"backend/internal/adapters/brokenrepo"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"errors"
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

	tokenService := domain.NewFakeTokenService(domain.WithToken("some-token"))

	sut := New(playerIdGenerator, playerRepository, tokenService)

	// Act
	result, err := sut.Execute("John Doe")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, stubbedPlayerId, result.id)
	assert.Equal(t, "some-token", result.token)

	createdPlayer, err := playerRepository.FindById(expectedPlayer.Id())
	require.NoError(t, err)
	assert.Equal(t, expectedPlayer, createdPlayer)
}

func Test_Create_player_with_broken_repository(t *testing.T) {
	// Arrange
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()
	playerRepository := brokenrepo.NewPlayerRepository()
	tokenService := domain.NewFakeTokenService()
	sut := New(playerIdGenerator, playerRepository, tokenService)

	// Act
	createdUserId, err := sut.Execute("some")

	// Assert
	assert.Empty(t, createdUserId)
	assert.NotNil(t, err)
}

func Test_Create_player_with_broken_token_service(t *testing.T) {
	// Arrange
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()
	playerRepository := inmemoryrepo.NewPlayerRepository()
	tokenService := domain.NewFakeTokenService(domain.WithError(errors.New("oops")))
	sut := New(playerIdGenerator, playerRepository, tokenService)

	// Act
	createdUserId, err := sut.Execute("some")

	// Assert
	assert.Empty(t, createdUserId)
	assert.NotNil(t, err)
}
