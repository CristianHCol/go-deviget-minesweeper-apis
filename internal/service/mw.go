package mw

import (
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infraestructure/common/cache"
	"github.com/CristianHCol/go-deviget-minesweeper-apis/internal/infraestructure/mw/repository"
)

// Minesweeper implements the mineswiperr contract
type Minesweeper struct {
	repo  repository.Service
	cache cache.Cache
}

// New creates a new instance for the minesweeper service
func New(minesweeperRepo repository.Service, servCache cache.Cache) *Minesweeper {
	return &Minesweeper{repo: minesweeperRepo, cache: servCache}
}
