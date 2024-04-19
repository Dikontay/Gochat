package repository

import (
	"database/sql"
	"gochat/internal/repository/user"
)

type Repository struct {
	User user.UserRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User: user.UserStorage{DB: db},
	}
}
