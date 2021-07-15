package domain

import (
	"testing"
)

func Test_New_Player(t *testing.T) {
	type args struct {
		id   PlayerId
		name string
	}
	tests := []struct {
		name string
		args args
		want Player
	}{
		{
			name: "Create a new player",
			args: args{
				id:   "some id",
				name: "some player",
			},
			want: Player{
				id:   "some id",
				name: "some player",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlayer(tt.args.id, tt.args.name)

			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
