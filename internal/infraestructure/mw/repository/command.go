package repository

import (
	"context"
)

// MinesWeeperCommand user command implementation
type MinesWeeperCommand struct{}

// NewMinesWeeperCommand initializes a new user command
func NewMinesWeeperCommand() *MinesWeeperCommand {
	return &MinesWeeperCommand{}
}

// CreateUser create a new user in the DB
func (c MinesWeeperCommand) CreateUser(ctx context.Context, userName string) (bool, error) {
	//TODO Save in redis
	return true, nil
}
