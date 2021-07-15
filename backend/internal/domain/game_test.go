package domain

import (
	"backend/internal/assertions"
	"testing"
)

func Test_first_player_joins_game(t *testing.T) {
	sut := NewGame()
	player := NewPlayer("some-id", "player-1")
	wantPlayers := []Player{player}

	sut.AddPlayer(player)

	assertions.DeepEqual(t, sut.Players(), wantPlayers)
}

func Test_second_player_joins_game(t *testing.T) {
	sut := NewGame()
	player1 := NewPlayer("1", "player-1")
	player2 := NewPlayer("2", "player-2")
	wantPlayers := []Player{player1, player2}

	sut.AddPlayer(player1)
	sut.AddPlayer(player2)

	assertions.DeepEqual(t, sut.Players(), wantPlayers)
}
