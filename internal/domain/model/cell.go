package model

// CellGrid
type CellGrid []Cell

// Cell object
type Cell struct {
	Mine        bool   `json:"mine"`
	Clicked     bool   `json:"clicked"`
	Flagged     bool   `json:"flagged"` // add a red flag in the cell
	Marked      bool   `json:"marked"`  // add a question mark
	Coordinates string `json:"coordinates"`
}
