package repository

// Service general minesweeper repository service, acts as a wrapper to avoid mulitple injection in upper layers.
type Service interface {
	CommandService
	QueryService
}
