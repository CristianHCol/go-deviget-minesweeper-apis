## Go Minesweeper apis

Contains all Apis for the Minesweeper game. Written in Golang and Redis.

## Requeriments
- For deploy locally, Redis migth be installed
- Use docker compose to setup your environment locally

## Installation
- Export environment variables source .env.dev or set it manually
- You could use go run cmd/mw/main.go 
- Or you could run make run command

## How play it?
- Create an user
```shell script
curl -X POST http://localhost:8090/api/v1/mw/user 
  -d '{
	"user_name": "Kriz"
}'
```

-  Create a game (If the user is not created system will return an USR_404)
- **user_name** the username above created
- **colums** Number of colums of game
- **rows** Number of rows of game
- **mines** Number of mines of game

```shell script
curl -X POST http://localhost:8090/api/v1/mw/game 
  -d '{
    "user_name": "Kriz",
    "colums": 10,
    "rows": 10,
    "mines": 5
}'
```

-- Start the game (If the user is not created system will return an USR_404)
- **game_name** the create game api return the name of the game
- **user_name** the username above created

```shell script
curl -X GET http://localhost:8090/api/v1/mw/game/{game_name}/{user_name}
```

-  Execute an action (If the user is not created system will return an USR_404)
for action_type you migth use: 'CLICK' or 'FLAG'
- **game_name** the create game api return the name of the game
- **user_name** the username above created
- **row** row to mark or flag
- **column** column to mark or flag

```shell script
curl -X POST http://localhost:8090/api/v1/mw/game/{game_name}/{user_name}/action 
  -d '{
    "action_type": "CLICK",
    "row": 4,
    "column": 1
}'
```

### This apis are deployed in Heroku
```
https://cristian-go-minesweeper-api.herokuapp.com/api/v1/mw/game/test/test
```