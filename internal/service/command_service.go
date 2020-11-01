package mw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
	dmmw "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
)

// CreateUser creates a new user in the system
func (s *Minesweeper) CreateUser(ctx context.Context, userName string) (bool, error) {
	userExist, err := s.cache.Get(ctx, userName)

	if err != nil || userExist != nil {
		fmt.Println("Username already exists: " + userName)
		return false, dmmw.ErrUserKeyDuplicated
	}

	user := []byte(userName)
	if err := s.cache.Set(ctx, userName, user, uint64(0)); err != nil {
		log.Fatal("error saving in redis")
		return false, dmmw.ErrInternalUser
	}
	fmt.Println("Username saved sucessfully")
	return true, nil
}

// CreateGame creates a new game in the system
func (s *Minesweeper) CreateGame(ctx context.Context, game *model.Game) (bool, error) {
	userExist, err := s.cache.Get(ctx, game.UserName)
	count := 0
	if err != nil || userExist == nil {
		fmt.Println("Username not exists: " + game.UserName)
		return false, dmmw.ErrUserNotFound
	}

	lastGame, err := s.cache.Get(ctx, game.UserName+"_LastGame")
	if err != nil {
		fmt.Println("Error getting last game")
		return false, dmmw.ErrInternalUser
	}
	if lastGame == nil {
		count = 1
	} else {
		byteToInt, _ := strconv.Atoi(string(lastGame))
		count = byteToInt + 1
		fmt.Println(count)
	}
	game.Name = fmt.Sprint(game.UserName+"Game", count)

	if game.Rows == 0 {
		game.Rows = model.DefaultRows
	}

	if game.Cols == 0 {
		game.Cols = model.DefaultCols
	}

	if game.Mines == 0 {
		game.Mines = model.DefaultMines
	}

	if game.Mines > model.MaxMines {
		game.Mines = model.MaxMines
	}

	if game.Rows > model.MaxRows {
		game.Rows = model.MaxRows
	}
	if game.Cols > model.MaxCols {
		game.Cols = model.MaxCols
	}

	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = (game.Cols * game.Rows)
	}

	gameBytes := new(bytes.Buffer)
	json.NewEncoder(gameBytes).Encode(game)
	GameToSave := []byte(gameBytes.Bytes())

	if err := s.cache.Set(ctx, game.Name, GameToSave, uint64(0)); err != nil {
		log.Fatal("error saving in redis")
		return false, dmmw.ErrInternalUser
	}

	if err := s.cache.Set(ctx, game.UserName+"_LastGame", []byte(strconv.Itoa(count)), uint64(0)); err != nil {
		log.Fatal("error saving lastgame in redis")
	}

	fmt.Println("Game saved sucessfully")
	return true, nil
}
