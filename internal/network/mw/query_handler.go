package mwhttp

import (
	"github.com/valyala/fasthttp"
)

// GetUserByID is the get user by id handler
func (h *MinesweeperHandler) GetUserByID(rctx *fasthttp.RequestCtx) {
	// rq := &mw.IDUserRequest{}
	// response := &common.BaseHTTPResponse{}
	// netcommon.SetCommonHeaders(rctx) // TOOD move this to the middleware

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()

	// id, err := strconv.Atoi(rctx.UserValue("id").(string))
	// if err != nil {
	// 	SetBadRequest(rctx)
	// 	return
	// }
	// rq.ID = uint64(id)
	// request := user.MapToIDUserRequest(rq)
	// rs, err := h.svc.GetUserById(ctx, request)
	// if err != nil {
	// 	if errors.Is(err, user.ErrorValidation) {
	// 		response = user.MapToUserResponse(nil, rs.Violations)
	// 	}
	// 	fmt.Println("[user Handler]: error getting user", err)
	// 	SetMetaAndError(response, err)
	// 	statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
	// 	rctx.SetBody(MarshalResponse(response))
	// 	rctx.SetStatusCode(statusCode)
	// 	return
	// }

	// response = user.MapToUserResponse(rs.Response, rs.Violations)
	// SetMeta(response)
	// response.Data = rs
	// rctx.SetBody(MarshalResponse(response))
	// rctx.SetStatusCode(fasthttp.StatusOK)
}
