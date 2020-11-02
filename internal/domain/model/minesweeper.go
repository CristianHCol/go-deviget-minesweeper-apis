package model

import (
	"context"
)

// MinesweeperServiceServer is the server API for UserService service.
type MinesweeperServiceServer interface {
	// CreateUser creates a new user in the system
	CreateUser(ctx context.Context, userName string) (bool, error)
	// CreateGame creates a new game in the system
	CreateGame(ctx context.Context, game *Game) (bool, string, error)
	// StartGame Start a new game
	StartGame(ctx context.Context, userName string, gameName string) (*Game, error)
	// ActionGame Action executed in the game while is playing
	ActionGame(ctx context.Context, userName string, gameName string, actionType string, row int, col int) (*Game, error)
}
