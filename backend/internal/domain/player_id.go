package domain

import "github.com/google/uuid"

type PlayerId interface {
	String() string
}

func ParsePlayerId(s string) (PlayerId, error) {
	value, err := uuid.Parse(s)
	if err != nil {
		return nil, err
	}
	return value, nil
}
