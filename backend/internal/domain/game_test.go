package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_First_player_joins_game(t *testing.T) {
	// arrange
	sut := NewGame()
	player := NewPlayer(NewFakePlayerId("some"), "player-1")
	expectedPlayers := []*Player{player}

	// act
	sut.AddPlayer(player)

	// assert
	assert.Equal(t, expectedPlayers, sut.Players())
}

func Test_Second_player_joins_game(t *testing.T) {
	// arrange
	sut := NewGame()
	player1 := NewPlayer(NewFakePlayerId("1"), "player-1")
	player2 := NewPlayer(NewFakePlayerId("2"), "player-2")
	expectedPlayers := []*Player{player1, player2}

	// act
	sut.AddPlayer(player1)
	sut.AddPlayer(player2)

	// assert
	assert.Equal(t, expectedPlayers, sut.Players())
}
