package token

import (
	"context"
	"database/sql"
	"gochat/internal/models"
	"time"
)

type TokenStorage struct {
	DB *sql.DB
}

type TokenRepo interface {
	New(userID int64, ttl time.Duration, scope string) (*models.Token, error)
	Insert(token *models.Token) error
	DeleteAllForUser(scope string, userID int64) error
}

func (s TokenStorage) New(userID int64, ttl time.Duration, scope string) (*models.Token, error) {
	token, err := models.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}
	err = s.Insert(token)
	return token, err
}

func (s TokenStorage) Insert(token *models.Token) error {
	query := `INSERT INTO tokens (hash, user_id, expiry, scope) VALUES ($1, $2, $3, $4)`
	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := s.DB.ExecContext(ctx, query, args...)
	return err
}

func (s TokenStorage) DeleteAllForUser(scope string, userID int64) error {
	query := `DELETE FROM tokens WHERE scope = $1 AND user_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := s.DB.ExecContext(ctx, query, scope, userID)
	return err
}
