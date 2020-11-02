package mwhttp

import (
	"context"
	"fmt"
	"time"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
	netcommon "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/common"

	"github.com/valyala/fasthttp"
)

// StartGame is the gstart game handler
func (h *MinesweeperHandler) StartGame(rctx *fasthttp.RequestCtx) {
	response := &common.BaseHTTPResponse{}
	netcommon.SetCommonHeaders(rctx)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	Name := rctx.UserValue("gamename").(string)
	UserName := rctx.UserValue("username").(string)

	if Name == "" || UserName == "" {
		SetBadRequest(rctx)
		return
	}

	rs, err := h.svc.StartGame(ctx, UserName, Name)
	if err != nil {
		fmt.Println("[start game Handler]: error starting game", err)
		SetMetaAndError(response, err)
		statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
		rctx.SetBody(MarshalResponse(response))
		rctx.SetStatusCode(statusCode)
		return
	}

	response = mw.MapToGameResponse(rs)
	SetMeta(response)
	response.Data = rs
	rctx.SetBody(MarshalResponse(response))
	rctx.SetStatusCode(fasthttp.StatusOK)
}
