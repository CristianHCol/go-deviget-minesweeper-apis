package model

// Action object
type Action struct {
	Type   string `json:"action_type"`
	Row    int    `json:"row"`
	Column int    `json:"column"`
}
