package repository

import (
	"context"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
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

// CreateGame create a new user in the DB
func (c MinesWeeperCommand) CreateGame(ctx context.Context, game model.Game) (bool, string, error) {
	//TODO Save in redis
	return true, "", nil
}

// ActionGame execute an action in the game
func (c MinesWeeperCommand) ActionGame(ctx context.Context, userName string, gameName string, actionType string, row int, col int) (*model.Game, error) {
	//TODO Save in redis
	response := &model.Game{}

	return response, nil
}
