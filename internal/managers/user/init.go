package user

import (
	"context"

	models "fe-sem4/internal/models/user"
)

type userStorer interface {
	CreateUser(ctx context.Context, user models.User) error
	FindUserByPhone(ctx context.Context, phone string) (*models.User, error)
	FindUserByID(ctx context.Context, id uint) (*models.User, error)
}

type Manager struct {
	userStorer userStorer
}

func NewManager(userStorer userStorer) *Manager {
	return &Manager{
		userStorer: userStorer,
	}
}
