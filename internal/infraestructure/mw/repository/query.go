package repository

import (
	"context"
)

// MinesWeeperQuery user query implementation
type MinesWeeperQuery struct{}

// NewMinesWeeperQuery initializes the user query
func NewMinesWeeperQuery() *MinesWeeperQuery {
	return &MinesWeeperQuery{}
}

// GetByUserName gets a user by the username
func (q *MinesWeeperQuery) GetByUserName(ctx context.Context, userName string) (bool, error) {
	return true, nil
}
