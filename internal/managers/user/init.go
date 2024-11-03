package user

import (
	"context"
	"fe-sem4/internal/models/user"
)

type userStorer interface {
	CreateUser(ctx context.Context, model user.User) error
	GetUserByPhone(ctx context.Context, phone string) (user.User, error)
}

type Manager struct {
	userStorer userStorer
}

func NewUserManager(userStorer userStorer) *Manager {
	return &Manager{userStorer: userStorer}
}
