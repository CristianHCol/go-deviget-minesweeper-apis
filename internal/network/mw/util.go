package mwhttp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
	netcommon "github.com/CristianHCol/go-deviget-minesweeper-apis/internal/network/common"
	"github.com/valyala/fasthttp"
)

var (
	service        = "minesweeper"
	serviceVersion = os.Getenv("API_VERSION")
)

// SetCommonHeaders set common heders for apis
func SetCommonHeaders(ctx *fasthttp.RequestCtx) {
	netcommon.SetCommonHeaders(ctx)
}

// SetBadRequest for users object
func SetBadRequest(ctx *fasthttp.RequestCtx) {
	netcommon.SetBadRequest(ctx, service, serviceVersion)
}

// SetMetaAndError sets the metadata and error object
func SetMetaAndError(response *common.BaseHTTPResponse, err error) {
	netcommon.SetMeta(response, service, serviceVersion)
	netcommon.SetError(response, err)
}

// SetMeta sets the response metadata
func SetMeta(response *common.BaseHTTPResponse) {
	netcommon.SetMeta(response, service, serviceVersion)
}

// SetError sets the error to the response object
func SetError(response *common.BaseHTTPResponse, err error) {
	netcommon.SetError(response, err)
}

// GetErrorHTTPStatusCode get the http status error code based in the service error code
func GetErrorHTTPStatusCode(statusCode string) int {
	return netcommon.GetErrorHTTPStatusCode(statusCode)
}

// MarshalResponse marshals the generic response object
func MarshalResponse(rs *common.BaseHTTPResponse) []byte {
	return netcommon.MarshalResponse(rs)
}

// GetIDFromRequest from request
func GetIDFromRequest(rctx *fasthttp.RequestCtx) (id int, err error) {
	id, err = strconv.Atoi(rctx.UserValue("id").(string))
	if err != nil {
		SetBadRequest(rctx)
		return 0, err
	}
	return id, nil
}

// SetCommonErrorResponse Build a generic error response
func SetCommonErrorResponse(err error, response *common.BaseHTTPResponse, violations map[string]string, rctx *fasthttp.RequestCtx, action string) {
	if err == mw.ErrorValidation {
		response = mw.MapToUserResponse(nil, violations)
	}
	fmt.Println("[user Handler]: error "+action+" Minesweeper", err)
	SetMetaAndError(response, err)
	statusCode := GetErrorHTTPStatusCode(response.Error.ErrorCode)
	rctx.SetBody(MarshalResponse(response))
	rctx.SetStatusCode(statusCode)
}

// SetCommonSucessResponse Build a generic sucess response
func SetCommonSucessResponse(response *common.BaseHTTPResponse, rctx *fasthttp.RequestCtx, httpStatusCode int) {
	SetMeta(response)
	rctx.SetBody(MarshalResponse(response))
	rctx.SetStatusCode(httpStatusCode)
}
