package user

import (
	"context"
	"fe-sem4/internal/models/session"
	"fe-sem4/internal/models/user"
)

type userStorer interface {
	CreateUser(ctx context.Context, model user.User) error
	GetUserByPhone(ctx context.Context, phone string) (user.User, error)
}

type sessionCreator interface {
	CreateSession(ctx context.Context, sess session.Session) error
}

type Manager struct {
	userStorer     userStorer
	sessionCreator sessionCreator
}

func NewUserManager(
	userStorer userStorer,
	sessionCreator sessionCreator,
) *Manager {
	return &Manager{
		userStorer:     userStorer,
		sessionCreator: sessionCreator,
	}
}
