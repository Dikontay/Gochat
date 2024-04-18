package user

import (
	"database/sql"
	"errors"
	"gochat/internal/models"
)

type UserStorage struct {
	DB *sql.DB
}

type UserRepo interface {
	CreateUser(user *models.User) error
	Get(id int64) (*models.User, error)
	Update(user *models.User) error
}

var ErrRecordNotFound = errors.New("record not found")

//ID        int       `json:"id"`
//Username  string    `json:"username"`
//Email     string    `json:"email"`
//Password  string    `json:"password"`
//CreatedAt time.Time `json:"-"`

func (s UserStorage) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	args := []any{user.Username, user.Email, user.Password}
	return s.DB.QueryRow(query, args).Scan(&user.ID, &user.Password)
}

func (m UserStorage) Get(id int64) (*models.User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `SELECT id,username, email, password, year FROM users WHERE id = $1`

	var user models.User
	err := m.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
		&user.Password,
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
