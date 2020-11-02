package repository

// MinesWeeperRepository is a wrapper for the command and query implementations injections are less in upper layers
type MinesWeeperRepository struct {
	*MinesWeeperCommand
	*MinesWeeperQuery
}

// NewMinesWeeperRepo initializes the instance for the repository
func NewMinesWeeperRepo() *MinesWeeperRepository {
	return &MinesWeeperRepository{&MinesWeeperCommand{}, &MinesWeeperQuery{}}
}
