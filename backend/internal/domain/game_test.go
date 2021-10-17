package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_First_player_joins_game(t *testing.T) {
	// Arrange
	sut := NewGame()
	player := NewPlayer(NewFakePlayerId("some"), "player-1")
	expectedPlayers := []*Player{player}

	// Act
	sut.AddPlayer(player)

	// Assert
	assert.Equal(t, expectedPlayers, sut.Players())
}

func Test_Second_player_joins_game(t *testing.T) {
	// Arrange
	sut := NewGame()
	player1 := NewPlayer(NewFakePlayerId("1"), "player-1")
	player2 := NewPlayer(NewFakePlayerId("2"), "player-2")
	expectedPlayers := []*Player{player1, player2}

	// Act
	sut.AddPlayer(player1)
	sut.AddPlayer(player2)

	// Assert
	assert.Equal(t, expectedPlayers, sut.Players())
}
