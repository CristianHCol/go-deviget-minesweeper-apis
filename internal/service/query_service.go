package mw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/service/common"

	dmmw "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
)

// StartGame start the game and return the grid
func (s *Minesweeper) StartGame(ctx context.Context, userName string, gameName string) (*model.Game, error) {
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

	if game.Status != "NEW" {
		fmt.Println("The game is already started: " + gameName)
		return nil, dmmw.ErrGameAlreadyStarted
	}

	response = common.CreateBoard(&game)
	response.Status = "PLAYING"
	response.StartedAt = time.Now()

	gameBytes := new(bytes.Buffer)
	json.NewEncoder(gameBytes).Encode(response)
	GameToSave := []byte(gameBytes.Bytes())

	if err := s.cache.Set(ctx, game.Name, GameToSave, uint64(0)); err != nil {
		log.Fatal("error saving in redis")
		return nil, dmmw.ErrInternalError
	}

	fmt.Println(response)
	return response, nil
}
