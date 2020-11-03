package netcommon

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/mw"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var (
	serviceName    = "svc_test"
	serviceVersion = "v1"
)

func TestSetCommonHeaders(t *testing.T) {
	request := &fasthttp.RequestCtx{}
	SetCommonHeaders(request)
	ct := request.Response.Header.Peek(fasthttp.HeaderContentType)
	assert.Equal(t, string(ct), "application/json; charset=UTF-8")
}

func TestSetBadRequest(t *testing.T) {
	request := &fasthttp.RequestCtx{}
	sample := &common.BaseHTTPResponse{}
	sample.Error = &common.ErrorResponse{ErrorCode: "BR_001", Message: "error reading request"}
	SetMeta(sample, serviceName, serviceVersion)
	SetBadRequest(request, serviceName, serviceVersion)
	rs := request.Response.Body()
	bt, _ := json.Marshal(sample)
	assert.Equal(t, bt, rs)
}

func TestSetMetaAndError(t *testing.T) {
	rs := &common.BaseHTTPResponse{}
	SetMetaAndError(rs, errors.New("sample error"), serviceName, serviceVersion)
	assert.NotNil(t, rs.Meta)
	assert.Equal(t, rs.Error.ErrorCode, mw.ErrInternalError)
}

func TestSetMeta(t *testing.T) {
	rs := &common.BaseHTTPResponse{}
	SetMeta(rs, serviceName, serviceVersion)
	assert.NotNil(t, rs.Meta)
}

func TestSetError(t *testing.T) {
	rs := &common.BaseHTTPResponse{}
	SetError(rs, nil)
	assert.Nil(t, rs.Error)
	SetError(rs, errors.New("sample error"))
	assert.Equal(t, rs.Error.ErrorCode, mw.ErrInternalError)
	SetError(rs, mw.ErrInternalError)
	assert.Error(t, mw.ErrInternalError, rs.Error.ErrorCode)
}

func TestGetErrorHTTPStatusCode(t *testing.T) {
	code1 := GetErrorHTTPStatusCode("USR_001")
	code2 := GetErrorHTTPStatusCode("NON_EXISTING_003")

	assert.Equal(t, http.StatusBadRequest, code1)
	assert.Equal(t, http.StatusInternalServerError, code2)
}

// MarshalResponse marshals the generic response object
func TestMarshalResponse(t *testing.T) {
	rs := &common.BaseHTTPResponse{}
	rs2 := MarshalResponse(rs)
	bt, _ := json.Marshal(rs)
	assert.Equal(t, rs2, bt)
}
