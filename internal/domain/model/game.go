package model

import (
	"time"
)

// Default values
const (
	DefaultRows  = 8
	DefaultCols  = 8
	DefaultMines = 4
	MaxMines     = 10
	MaxRows      = 30
	MaxCols      = 30
)

// Game object
type Game struct {
	Name      string        `json:"name"`
	UserName  string        `json:"user_name"`
	Rows      int           `json:"rows"`
	Cols      int           `json:"cols"`
	Mines     int           `json:"mines"`
	Status    string        `json:"status"` //NEW, PLAYING, WON, GAME_OVER
	Grid      []CellGrid    `json:"grid,omitempty"`
	Cliks     int           `json:"clicks"`
	Username  string        `json:"username"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	StartedAt time.Time     `json:"started_at"`
	TimeSpent time.Duration `json:"time_spent"`
	Points    float32       `json:"points,omitempty"`
}
