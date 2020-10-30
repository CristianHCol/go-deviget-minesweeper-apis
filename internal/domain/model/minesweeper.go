package model

import "context"

// MinesweeperServiceServer is the server API for UserService service.
type MinesweeperServiceServer interface {
	// CreateUser creates a new user in the system
	CreateUser(ctx context.Context, userName string) (bool, error)
}
