package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

)

// MinesWeeperQuery user query implementation
type MinesWeeperQuery struct{}

// NewMinesWeeperQueryQuery initializes the user query
func NewMinesWeeperQueryQuery() *MinesWeeperQuery {
	return &MinesWeeperQuery{}
}

// GetByUserName gets a user by the username
func (q *MinesWeeperQuery) GetByUserName(ctx context.Context, userName string) (userExists bool, error) {
	return true, nil
}
