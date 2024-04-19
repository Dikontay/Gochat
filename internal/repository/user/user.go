package user

import (
	"context"
	"database/sql"
	"errors"
	"gochat/internal/models"
	"time"
)

type UserStorage struct {
	DB *sql.DB
}

type UserRepo interface {
	CreateUser(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int64) error
}

var ErrRecordNotFound = errors.New("record not found")
var ErrDuplicateEmail = errors.New("duplicate email")

//ID        int       `json:"id"`
//Username  string    `json:"username"`
//Email     string    `json:"email"`
//Password  string    `json:"password"`
//CreatedAt time.Time `json:"-"`

func (s UserStorage) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash, activated) VALUES ($1, $2, $3) RETURNING id, created_at`

	args := []any{user.Username, user.Email, user.Password}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := s.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return s.DB.QueryRow(query, args).Scan(&user.ID, &user.CreatedAt)
}

func (m UserStorage) GetByEmail(email string) (*models.User, error) {
	query := `
SELECT id, created_at, name, email, password_hash, activated, version FROM users
WHERE email = $1`
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&user.ID,
		&user.CreatedAt, &user.Username, &user.Email, &user.Password.Hash, &user.Activated,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (s UserStorage) Update(user *models.User) error {
	query := `UPDATE users 
	SET username = $1, email = $2, password = $3
	WHERE id = $4`

	args := []any{
		user.Username, user.Email, user.Password,
	}

	return s.DB.QueryRow(query, args).Err()
}

func (s UserStorage) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE FROM users WHERE id = $1`

	result, err := s.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
