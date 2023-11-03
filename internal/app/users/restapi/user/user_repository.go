package user

import (
	"github.com/mcarello/user_service/internal/app/users/database"
)

type Repository interface {
	Get() string
}

type repository struct {
	*database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

func (r *repository) Get() string {
	return "Hello From Repository"
}
