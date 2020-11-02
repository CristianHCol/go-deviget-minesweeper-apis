## Go Minesweeper apis

Contains all Apis for the Minesweeper game. Written in Golang and Redis.

## TODO requeriments
- For deploy locally, Redis migth be installed
- Use docker compose to setup your environment locally

## TODO installation
- You could use go run cmd/mw/main.go 
- Or you could run make run command

## TODO How play it?
- Create an user
```shell script
curl -X POST \
  http://localhost:8090/api/v1/mw/users \
  -d '{
	"user_name": "Kriz"
}'
```

-  Create a game (If the user is not created system will return an USR_404)
```shell script
curl -X POST \
  http://localhost:8090/api/v1/mw/users \
  -d '{
    "user_name": "Kriz",
    "colums": 10,
    "rows": 10,
    "mines": 5
}'
```

-- Start the game (If the user is not created system will return an USR_404)
```shell script
curl -X GET \
  http://localhost:8090/api/v1/mw/game/{game_name}/{user_name} \
```

-  Execute an action (If the user is not created system will return an USR_404)
for action_type you migth use: 'CLICK' or 'FLAG'

```shell script
curl -X POST \
  http://localhost:8090/api/v1/mw/game/KrizGame1/Kriz/action \
  -d '{
    "action_type": "CLICK",
    "row": 4,
    "column": 1
}'
```

### This api is deployed in heroku
```
https://cristian-go-minesweeper-api.herokuapp.com/api/v1/mw/user
```