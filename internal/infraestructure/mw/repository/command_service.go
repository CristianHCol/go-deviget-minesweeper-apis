package repository

import (
	"context"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// CommandService operation that writes/deletes something in the database
type CommandService interface {
	// CreateUser a new user
	CreateUser(ctx context.Context, userName string) (bool, error)
	// CreateGame creates a new game in the system
	CreateGame(ctx context.Context, game model.Game) (bool, error)
}
