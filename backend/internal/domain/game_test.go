package domain

import (
	"reflect"
	"testing"
)

func Test_Game_Add_Player(t *testing.T) {
	type fields struct {
		players []Player
	}
	type args struct {
		p Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Game
	}{
		{
			name: "Add player to an empty game",
			fields: fields{
				players: nil,
			},
			args: args{
				p: Player{
					name: "some player",
				},
			},
			want: &Game{
				players: []Player{
					{name: "some player"},
				},
			},
		},
		{
			name: "Add player to a non-empty game",
			fields: fields{
				players: []Player{
					{
						name: "player1",
					},
				},
			},
			args: args{
				p: Player{
					name: "player2",
				},
			},
			want: &Game{
				players: []Player{
					{name: "player1"},
					{name: "player2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := &Game{
				players: tt.fields.players,
			}

			sut.addPlayer(tt.args.p)

			if !reflect.DeepEqual(sut, tt.want) {
				t.Errorf("Got %v, want %v", sut, tt.want)
			}
		})
	}
}

func Test_New_Game(t *testing.T) {
	type args struct {
		id GameId
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "Create a new game",
			args: args{
				id: "some",
			},
			want: Game{
				id:      "some",
				players: []Player{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
