package main

import (
	"log"
	"os"

	netcommon "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/common"
	server "github.com/CristianHCol/lm-http/pkg/server"
	"github.com/joho/godotenv"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infrastructure/common/cache"
	mwrepo "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infrastructure/mw/repository"
	mwhttp "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/mw/http"
	mwrsvc "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/service/mw"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	netcommon.LoadConfig()
	mwRepo := mwrepo.NewMinesWeeperRepo()
	cacheSvc := cache.NewRedisInstance()
	mwSvc := mwrsvc.New(mwrepo, cacheSvc)
	mwHandler := mwhttp.NewMinesWeeperHandler(mwSvc)
	sv := &server.Server{}
	sv.AddHandler(netcommon.MinesWeeperBasePath, netcommon.Post, mwHandler.CreateUser)

	go sv.Start(os.Getenv("HTTP_PORT"))
}
