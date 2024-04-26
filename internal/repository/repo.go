package repository

import (
	"database/sql"
	"gochat/internal/repository/token"
	"gochat/internal/repository/user"
)

type Repository struct {
	User  user.UserRepo
	Token token.TokenRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:  user.UserStorage{DB: db},
		Token: token.TokenStorage{DB: db},
	}
}
