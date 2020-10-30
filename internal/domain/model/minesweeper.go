package model

import "context"

// MinesweeperServiceServer is the server API for UserService service.
type MinesweeperServiceServer interface {
	// CreateUser creates a new user in the system
	CreateUser(context.Context, userName string) (userStatus bool, error)
}
