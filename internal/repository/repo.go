package repository

import "gochat/internal/repository/user"

type Repository struct {
	User *user.UserRepo
}
