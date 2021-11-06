package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Parse_UUID_player_id(t *testing.T) {
	idString := "123e4567-e89b-12d3-a456-426614174000"

	id, err := ParsePlayerId(idString)

	assert.NoError(t, err)
	assert.Equal(t, idString, id.String())
}

func Test_Parse_invalid_UUID_player_id(t *testing.T) {
	idString := "not-a-uuid"

	_, err := ParsePlayerId(idString)

	assert.NotNil(t, err)
}
