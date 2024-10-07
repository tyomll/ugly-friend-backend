package repository

import (
	"context"
	"fmt"
	"ugly-friend/models"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) CreateUser(user *models.CreateUserReq) error {
	initialDebts := 0

	// Ensure that username is unique
	exists, err := r.userExists(user.Username)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("user_exists")
	}

	query := squirrel.Insert("users").
		Columns("username", "password", "card_numbers", "total_debts").
		Values(user.Username, user.Password, user.CardNumbers, initialDebts).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to convert query to sql: %w", err)
	}
	_, err = r.db.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *Repository) userExists(username string) (bool, error) {
	var count int
	err := r.db.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE username=$1", username).
		Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}
	return count > 0, nil
}
