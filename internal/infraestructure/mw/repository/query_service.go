package repository

import (
	"context"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// QueryService Query service contract
type QueryService interface {
	// StartGame start a new game by user
	StartGame(ctx context.Context, userName string, gameName string) (*model.Game, error)
}
