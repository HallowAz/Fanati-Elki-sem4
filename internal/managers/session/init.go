package session

import (
	"context"

	sessionmodels "fe-sem4/internal/models/session"
	usermodels "fe-sem4/internal/models/user"
)

type sessionStorer interface {
	Create(cookie *sessionmodels.Cookie) error
	Check(sessionToken string) (*sessionmodels.Cookie, error)
	Delete(cookie *sessionmodels.Cookie) error
	Expire(cookie *sessionmodels.Cookie) error
}

type userStorer interface {
	CreateUser(ctx context.Context, user usermodels.User) error
	FindUserByPhone(ctx context.Context, phone string) (*usermodels.User, error)
	FindUserByID(ctx context.Context, id uint) (*usermodels.User, error)
}

type Manager struct {
	sessionStorer sessionStorer
	userStorer    userStorer
}

func NewManager(sessionStorer sessionStorer, userStorer userStorer) *Manager {
	return &Manager{
		userStorer:    userStorer,
		sessionStorer: sessionStorer,
	}
}
