package mwhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"

	"github.com/valyala/fasthttp"
)

// CreateUser - create user handler
func (h *MinesweeperHandler) CreateUser(rctx *fasthttp.RequestCtx) {
	rq := &mw.CreateUserRequest{}
	response := &common.BaseHTTPResponse{}
	SetCommonHeaders(rctx)

	if err := json.Unmarshal(rctx.PostBody(), rq); err != nil {
		SetBadRequest(rctx)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	request := mw.MapToUserName(rq)
	rs, err := h.svc.CreateUser(ctx, request)
	if err != nil {
		fmt.Println("[create user Handler]: error creating User", err)
		SetMetaAndError(response, err)
		statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
		rctx.SetBody(MarshalResponse(response))
		rctx.SetStatusCode(statusCode)
		return
	}

	response = mw.MapToCreateResponse(rs, "user")
	SetCommonSucessResponse(response, rctx, fasthttp.StatusCreated)
}

// CreateGame - create game handler
func (h *MinesweeperHandler) CreateGame(rctx *fasthttp.RequestCtx) {
	rq := &mw.CreateGameRequest{}
	response := &common.BaseHTTPResponse{}
	SetCommonHeaders(rctx)

	if err := json.Unmarshal(rctx.PostBody(), rq); err != nil {
		SetBadRequest(rctx)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	request := mw.MapToGame(rq)
	rs, err := h.svc.CreateGame(ctx, request)
	if err != nil {
		fmt.Println("[create game Handler]: error Creation Game", err)
		SetMetaAndError(response, err)
		statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
		rctx.SetBody(MarshalResponse(response))
		rctx.SetStatusCode(statusCode)
		return
	}

	response = mw.MapToCreateResponse(rs, "game")
	SetCommonSucessResponse(response, rctx, fasthttp.StatusCreated)
}
