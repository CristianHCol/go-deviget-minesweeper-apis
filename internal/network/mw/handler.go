package mwhttp

import (
	mwmodel "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// MinesweeperHandler mw handler
type MinesweeperHandler struct {
	svc mwmodel.MinesweeperServiceServer
}

// NewMinesWeeperHandler intializes the user handler
func NewMinesWeeperHandler(service mwmodel.MinesweeperServiceServer) *MinesweeperHandler {
	return &MinesweeperHandler{svc: service}
}
