package repository

import (
	"context"
)

// QueryService Query service contract
type QueryService interface {
	// GetByUserName fetches one user from storage
	GetByUserName(ctx context.Context, userName string) (bool, error)
}
