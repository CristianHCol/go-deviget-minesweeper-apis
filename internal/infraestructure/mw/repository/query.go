package repository

import (
	"context"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// MinesWeeperQuery user query implementation
type MinesWeeperQuery struct{}

// NewMinesWeeperQuery initializes the user query
func NewMinesWeeperQuery() *MinesWeeperQuery {
	return &MinesWeeperQuery{}
}

// StartGame gstart a game and return board
func (q *MinesWeeperQuery) StartGame(ctx context.Context, userName string, gameName string) (*model.Game, error) {
	response := &model.Game{}

	return response, nil
}
