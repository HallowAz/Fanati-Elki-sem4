package repository

import (
	"context"
	"errors"
	"fe-sem4/internal/models/domain_error"
	"fmt"

	"fe-sem4/internal/models/user"
	"fe-sem4/internal/repository/db"
	user_db "fe-sem4/internal/repository/internal/user"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db db.TxCommitter
}

func NewUserRepo(db db.TxCommitter) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(ctx context.Context, model user.User) error {
	err := u.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		return user_db.CreateUser(ctx, tx, user_db.NewUserRow(model))
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (u *UserRepo) GetUserByID(ctx context.Context, id uint32) (user.User, error) {
	var (
		row user_db.UserRow
		err error
	)

	err = u.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		row, err = user_db.GetUserByID(ctx, tx, id)
		return err
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.User{}, domain_error.ErrUserNotFound
		}

		return user.User{}, fmt.Errorf("failed to get user by id: %w", err)
	}

	return row.ToModel(), nil
}

func (u *UserRepo) GetUserByPhone(ctx context.Context, phone string) (user.User, error) {
	var (
		row user_db.UserRow
		err error
	)

	err = u.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		row, err = user_db.GetUserByPhone(ctx, tx, phone)
		return err
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.User{}, domain_error.ErrUserNotFound
		}

		return user.User{}, fmt.Errorf("failed to get user by phone: %w", err)
	}

	return row.ToModel(), nil
}
