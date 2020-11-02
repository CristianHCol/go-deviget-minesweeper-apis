package mw

import (
	"testing"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestMapToCreateResponset(t *testing.T) {
	rs := MapToCreateResponse(true, "game", "game1")
	assert.Equal(t, rs.Data, "game game1 Created sucessfully")
}

func TestMapToGameResponse(t *testing.T) {
	game := &model.Game{}
	game.Name = "Game1"
	rs := MapToGameResponse(game)
	assert.Equal(t, rs.Data, game)
}
