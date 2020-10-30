package repository

import (
	"context"
)

// CommandService operation that writes/deletes something in the database
type CommandService interface {
	// Create a new user
	CreateUser(ctx context.Context, user string) (userStatus bool, error)
}
