package netcommon

import (
	"fmt"
	"os"
)

// http verbs
const (
	Post   = "POST"
	Get    = "GET"
	Put    = "PUT"
	Delete = "DELETE"
	Patch  = "PATCH"
	Option = "OPTION"
)

// APIVersion api version
var APIVersion = "v1"

// BasePathAPI base path
var BasePathAPI = "/api/v1"

// MinesWeeperBasePath game base path
var MinesWeeperBasePath = "/api/v1/mw"

// LoadConfig update default values
func LoadConfig() {
	APIVersion = os.Getenv("API_VERSION")
	BasePathAPI = "/api/" + APIVersion
	MinesWeeperBasePath = fmt.Sprintf("%s/mw", BasePathAPI)
}
