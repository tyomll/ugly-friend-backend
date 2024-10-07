package repository

import (
	"context"
	"fmt"
	"ugly-friend/models"

	"github.com/Masterminds/squirrel"
)

func (r *Repository) CreateUser(user *models.CreateUserReq) error {
	query := squirrel.Insert("users").
		Columns("username", "password", "card_numbers").
		Values(user.Username, user.Password, user.CardNumbers).
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
