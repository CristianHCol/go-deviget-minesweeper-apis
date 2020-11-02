package mw

import (
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// CreateUserRequest request for create user
type CreateUserRequest struct {
	UserName string `json:"user_name,omitempty"`
}

// CreateGameRequest request for create game
type CreateGameRequest struct {
	UserName string `json:"user_name,omitempty"`
	Columns  int    `json:"colums,omitempty"`
	Rows     int    `json:"rows,omitempty"`
	Mines    int    `json:"mines,omitempty"`
}

// ActionRequest request for create game
type ActionRequest struct {
	ActionType string `json:"action_type,omitempty"`
	Row        int    `json:"row,omitempty"`
	Colum      int    `json:"column,omitempty"`
}

// MapToUserName in request maps http request into user request
func MapToUserName(request *CreateUserRequest) string {
	return request.UserName
}

// MapToGame in request maps http request into user request
func MapToGame(request *CreateGameRequest) *model.Game {
	return &model.Game{
		UserName:  request.UserName,
		Name:      "",
		Rows:      request.Rows,
		Cols:      request.Columns,
		Mines:     request.Mines,
		Status:    "NEW",
		Grid:      nil,
		Cliks:     0,
		Points:    0,
		CreatedAt: time.Now(),
		StartedAt: time.Now(),
		TimeSpent: 0,
	}
}

// MapToAction in request maps http request into user request
func MapToAction(request *ActionRequest) *model.Action {
	return &model.Action{
		Type:   request.ActionType,
		Row:    request.Row,
		Column: request.Colum,
	}
}
