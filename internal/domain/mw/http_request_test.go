package mw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToMapToUserName maps http request into user request
func TestToMapToUserName(t *testing.T) {
	cur := &CreateUserRequest{UserName: "test"}
	usreName := MapToUserName(cur)
	assert.NotNil(t, cur.UserName, usreName)
	assert.Equal(t, usreName, cur.UserName)
}

// TestMapToMapToGame maps http request into user request
func TestMapToMapToGame(t *testing.T) {
	cur := &CreateGameRequest{UserName: "test-update"}
	gameReq := MapToGame(cur)
	assert.NotNil(t, gameReq, gameReq.Username)
	assert.Equal(t, gameReq.UserName, cur.UserName)
}

// MapMapToIDUserRequest maps http request into user request
func TestToMapToAction(t *testing.T) {
	cur := &ActionRequest{Row: 33}
	actionReq := MapToAction(cur)
	assert.NotNil(t, actionReq, actionReq.Row)
	assert.Equal(t, actionReq.Row, cur.Row)
}
