package repository

import (
	"context"
)

// CommandService operation that writes/deletes something in the database
type CommandService interface {
	// CreateUser a new user
	CreateUser(ctx context.Context, userName string) (bool, error)
}
