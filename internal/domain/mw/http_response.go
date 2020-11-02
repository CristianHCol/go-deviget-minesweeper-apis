package mw

import (
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/common"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/domain/model"
)

// MapToCreateResponse converts the service object into the http response object
func MapToCreateResponse(entityCreated bool, entityName string, entityValue string) *common.BaseHTTPResponse {
	rs := &common.BaseHTTPResponse{}
	if entityCreated {
		rs.Data = entityName + " " + entityValue + " Created sucessfully"
	}
	return rs
}

// MapToGameResponse converts the service object into the http response object
func MapToGameResponse(game *model.Game) *common.BaseHTTPResponse {
	rs := &common.BaseHTTPResponse{}
	if game != nil {
		rs.Data = game
	}
	return rs
}
