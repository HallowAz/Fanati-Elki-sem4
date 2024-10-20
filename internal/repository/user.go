package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "fe-sem4/internal/models/user"
	repository "fe-sem4/internal/repository/internal/user"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user models.User) error {
	err := repository.CreateUser(ctx, ur.DB, &user)
	if err != nil {
		return fmt.Errorf("failed to create form: %w", err)
	}

	return nil
}
