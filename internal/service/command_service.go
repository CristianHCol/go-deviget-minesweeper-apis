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
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/service/common"
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
		return false, dmmw.ErrInternalError
	}
	fmt.Println("Username saved sucessfully")
	return true, nil
}

// CreateGame creates a new game in the system
func (s *Minesweeper) CreateGame(ctx context.Context, game *model.Game) (bool, string, error) {
	userExist, err := s.cache.Get(ctx, game.UserName)
	count := 0
	if err != nil || userExist == nil {
		fmt.Println("Username not exists: " + game.UserName)
		return false, "", dmmw.ErrUserNotFound
	}

	lastGame, err := s.cache.Get(ctx, game.UserName+"_LastGame")
	if err != nil {
		fmt.Println("Error getting last game")
		return false, "", dmmw.ErrInternalError
	}
	if lastGame == nil {
		count = 1
	} else {
		byteToInt, _ := strconv.Atoi(string(lastGame))
		count = byteToInt + 1
	}
	game.Name = fmt.Sprint(game.UserName+"Game", count)
	game.Status = "NEW"

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
		return false, "", dmmw.ErrInternalError
	}

	if err := s.cache.Set(ctx, game.UserName+"_LastGame", []byte(strconv.Itoa(count)), uint64(0)); err != nil {
		log.Fatal("error saving lastgame in redis")
	}

	fmt.Println("Game saved sucessfully")
	return true, game.Name, nil
}

// ActionGame execute an action in the game
func (s *Minesweeper) ActionGame(ctx context.Context, userName string, gameName string, actionType string, row int, col int) (*model.Game, error) {
	response := &model.Game{}

	gameExist, err := s.cache.Get(ctx, gameName)
	if err != nil || gameExist == nil {
		fmt.Println("game do not exists: " + gameName)
		return nil, dmmw.ErrGameNotFound
	}
	var game model.Game
	err = json.Unmarshal([]byte(gameExist), &game)
	if err != nil {
		return nil, dmmw.ErrInternalError
	}

	if game.UserName != userName {
		fmt.Println("Invalid username for this game: " + userName)
		return nil, dmmw.ErrGamePermissionDenied
	}
	response = &game
	if game.Status == "GAME_OVER" {
		fmt.Println("The game is OVER: " + gameName)
		return response, dmmw.ErrGameIsOver
	}

	if game.Status == "WON" {
		fmt.Println("The game is WON: " + gameName)
		return response, dmmw.ErrGameIsWon
	}

	if game.Status == "PLAYING" {
		if actionType == "CLICK" {
			if err := common.Click(&game, row, col); err != nil {
				return nil, err
			}
		}

		if actionType == "FLAG" {
			if err := common.Flag(&game, row, col); err != nil {
				return nil, err
			}
		}
		response = &game
		gameBytes := new(bytes.Buffer)
		json.NewEncoder(gameBytes).Encode(response)
		GameToSave := []byte(gameBytes.Bytes())

		if err := s.cache.Set(ctx, game.Name, GameToSave, uint64(0)); err != nil {
			log.Fatal("error saving in redis")
			return nil, dmmw.ErrInternalError
		}

		fmt.Println(response)

		if game.Status == "GAME_OVER" {
			fmt.Println("The game is OVER: " + gameName)
			return response, dmmw.ErrGameIsOver
		}

		if game.Status == "WON" {
			fmt.Println("The game is WON: " + gameName)
			return response, dmmw.ErrGameIsWon
		}
	} else {
		fmt.Println("Unknown game status: " + game.Status)
		return response, dmmw.ErrInternalError
	}

	return response, nil
}
