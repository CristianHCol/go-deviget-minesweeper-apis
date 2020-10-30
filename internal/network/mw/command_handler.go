package mwhttp

import (
	"github.com/valyala/fasthttp"
)

// CreateUser is the create user handler
func (h *MinesweeperHandler) CreateUser(rctx *fasthttp.RequestCtx) {
	// rq := &mw.CreateUserRequest{}
	// response := &common.BaseHTTPResponse{}
	// SetCommonHeaders(rctx) // TOOD move this to the middleware

	// if err := json.Unmarshal(rctx.PostBody(), rq); err != nil {
	// 	SetBadRequest(rctx)
	// 	return
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()

	// request := mw.MapToUserRequest(rq)
	// rs, err := h.svc.CreateUser(ctx, request)
	// if err != nil {
	// 	if errors.Is(err, user.ErrorValidation) {
	// 		response = mw.MapToUserResponse(rs.Response, rs.Violations)
	// 	}
	// 	fmt.Println("[user Handler]: error Creation User", err)
	// 	SetMetaAndError(response, err)
	// 	statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
	// 	rctx.SetBody(MarshalResponse(response))
	// 	rctx.SetStatusCode(statusCode)
	// 	return
	// }

	// response = user.MapToUserResponse(rs.Response, rs.Violations)
	// SetCommonSucessResponse(response, rctx, fasthttp.StatusCreated)
}
