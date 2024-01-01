package repository

type repository struct {
}

// NewRepository
func NewRepository() RepositoryInterface {
	return &repository{
	}
}
