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
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (ur *UserRepo) FindUserByPhone(ctx context.Context, phone string) (*models.User, error) {
	user, err := repository.FindUserByPhone(ctx, ur.DB, phone)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by phone: %w", err)
	}

	return user, nil
}

func (ur *UserRepo) FindUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := repository.FindUserByID(ctx, ur.DB, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return user, nil
}
