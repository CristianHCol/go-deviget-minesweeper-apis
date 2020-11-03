package main

import (
	"log"
	"os"

	netcommon "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/common"
	server "github.com/CristianHCol/go-deviget-minesweeper-apis/server"
	"github.com/joho/godotenv"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infraestructure/common/cache"
	mwrepo "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infraestructure/mw/repository"
	mwhttp "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/mw"
	mwrsvc "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	netcommon.LoadConfig()
	mwRepo := mwrepo.NewMinesWeeperRepo()
	cacheSvc := cache.NewRedisInstance()
	mwSvc := mwrsvc.New(mwRepo, cacheSvc)
	mwHandler := mwhttp.NewMinesWeeperHandler(mwSvc)
	sv := &server.Server{}
	sv.AddHandler(netcommon.MinesWeeperBasePath+"/user", netcommon.Post, mwHandler.CreateUser)
	sv.AddHandler(netcommon.MinesWeeperBasePath+"/game", netcommon.Post, mwHandler.CreateGame)
	sv.AddHandler(netcommon.MinesWeeperBasePath+"/game/{gamename}/{username}", netcommon.Get, mwHandler.StartGame)
	sv.AddHandler(netcommon.MinesWeeperBasePath+"/game/{gamename}/{username}/action", netcommon.Post, mwHandler.ActionGame)

	sv.Start(os.Getenv("HTTP_PORT"))

}
