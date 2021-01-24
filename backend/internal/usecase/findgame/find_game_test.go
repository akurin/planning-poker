package findgame

import (
	"backend/internal/domain"
	"reflect"
	"testing"
)

func Test_findGame_Execute(t *testing.T) {
	type args struct {
		id domain.GameId
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Player
		wantErr bool
	}{
		{
			name:    "Find game",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := findGame{}
			got, err := f.Execute(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
