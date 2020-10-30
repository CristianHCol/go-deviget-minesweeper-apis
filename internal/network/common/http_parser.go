package netcommon

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
	"github.com/valyala/fasthttp"
)

var (
	badRequestMsg  = "error reading request"
	badRequestCode = "BR_001"
)

// SetCommonHeaders set common heders for apis
func SetCommonHeaders(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")
	// TODO set headers for cors
}

// SetBadRequest sets a bad request response
func SetBadRequest(ctx *fasthttp.RequestCtx, service string, serviceVersion string) {
	rs := &common.BaseHTTPResponse{}
	SetMeta(rs, service, serviceVersion)
	rs.Error = &common.ErrorResponse{ErrorCode: badRequestCode, Message: badRequestMsg}
	// omited error because input struct is controlled
	bt, _ := json.Marshal(rs)

	ctx.SetBody(bt)
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

// SetMetaAndError sets the metadata and error object
func SetMetaAndError(response *common.BaseHTTPResponse, err error, service string, serviceVersion string) {
	SetMeta(response, service, serviceVersion)
	SetError(response, err)
}

// SetMeta sets the response metadata
func SetMeta(response *common.BaseHTTPResponse, service string, serviceVersion string) {
	if response.Meta == nil {
		response.Meta = &common.Meta{}
	}

	response.Meta.Service = service
	response.Meta.Version = serviceVersion
}

// SetError sets the error to the response object
func SetError(response *common.BaseHTTPResponse, err error) {
	if err == nil {
		return
	}

	if response.Error == nil {
		response.Error = &common.ErrorResponse{}
	}

	if code, ok := mw.ServiceErrorsCode[err]; !ok {
		response.Error.ErrorCode = mw.UserInternalError
	} else {
		response.Error.ErrorCode = code
	}

	response.Error.Message = parseErrorMessage(err.Error())
}

// GetErrorHTTPStatusCode get the http status error code based in the service error code
func GetErrorHTTPStatusCode(statusCode string) int {
	code, ok := mw.HTTPErrorsCode[statusCode]
	if !ok {
		return http.StatusInternalServerError
	}

	return code
}

// MarshalResponse marshals the generic response object
func MarshalResponse(rs *common.BaseHTTPResponse) []byte {
	// error omited since the input is controlled
	bt, _ := json.Marshal(rs)
	return bt
}

func parseErrorMessage(message string) string {
	idx := strings.Index(message, "desc =")
	if idx > 0 {
		return message[idx+7:]
	}

	return message
}
