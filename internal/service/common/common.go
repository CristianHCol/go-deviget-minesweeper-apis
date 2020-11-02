package common

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// CellGrid create the game

// CreateBoard create the game
func CreateBoard(game *model.Game) *model.Game {
	cellsNumber := game.Cols * game.Rows
	cells := make(model.CellGrid, cellsNumber)

	// set mines ramdon
	i := 0
	for i < game.Mines {
		index := rand.Intn(cellsNumber)
		if !cells[index].Mine {
			cells[index].Mine = true
			i++
		}
	}

	game.Grid = make([]model.CellGrid, game.Rows)
	for row := range game.Grid {
		game.Grid[row] = cells[(game.Cols * row):(game.Cols * (row + 1))]
	}

	for row, col := range game.Grid {
		for c := range col {
			game.Grid[row][c].Coordinates = fmt.Sprintf("row: %d, col:%d", row, c)
		}
	}
	return game
}

// Click execute click
func Click(game *model.Game, i, j int) error {
	if game.Grid[i][j].Clicked {
		return errors.New("Already click this field")
	}
	game.Grid[i][j].Clicked = true
	if game.Grid[i][j].Mine {
		game.Status = "GAME_OVER"
		return nil
	}
	game.Cliks++
	if checkWon(game) {
		game.Status = "WON"
		game.TimeSpent = time.Now().Sub(game.StartedAt)
	}
	return nil
}

// Flag execute flag
func Flag(game *model.Game, i, j int) error {
	if game.Grid[i][j].Clicked {
		return errors.New("Already click this field")
	}

	if game.Grid[i][j].Flagged {
		return errors.New("Already flag this field")
	}
	return nil
}

func checkWon(game *model.Game) bool {
	started := game.Status == "PLAYING"
	return game.Cliks == ((game.Rows*game.Cols)-game.Mines) && started
}
