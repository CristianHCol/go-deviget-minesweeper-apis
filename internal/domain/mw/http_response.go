package mw

import (
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
)

// MapToCreateResponse converts the service object into the http response object
func MapToCreateResponse(entityCreated bool, entityName string) *common.BaseHTTPResponse {
	rs := &common.BaseHTTPResponse{}
	if entityCreated {
		rs.Data = entityName + " Created sucessfully"
	}
	return rs
}
