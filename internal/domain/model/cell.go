package model

//CellGrid
type CellGrid []Cell

// Cell object
type Cell struct {
	Mine        bool   `json:"mine"`
	Clicked     bool   `json:"clicked"`
	Flagged     bool   `json:"flagged"`
	Marked      bool   `json:"marked"`
	Coordinates string `json:"coordinates"`
}
