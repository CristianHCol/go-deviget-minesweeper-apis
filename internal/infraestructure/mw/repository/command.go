package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

)

// MinesWeeperCommand user command implementation
type MinesWeeperCommand struct{}

// NewMinesWeeperCommand initializes a new user command
func NewMinesWeeperCommand() *MinesWeeperCommand {
	return &MinesWeeperCommand{}
}

// Create creates a new user in the DB
func (c MinesWeeperCommand) Create(ctx context.Context, user string) (userStatus bool, error) {
	//TODO Save in redis
	return true, nil
}

