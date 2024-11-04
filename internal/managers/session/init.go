package session

import (
	"context"
	"fe-sem4/internal/models/session"
	"fe-sem4/internal/models/user"
)

type sessionRepo interface {
	GetSession(ctx context.Context, key string) (session.Session, error)
	CreateSession(ctx context.Context, sess session.Session) error
}

type userGetter interface {
	GetUserByPhone(ctx context.Context, phone string) (user.User, error)
}

type Manager struct {
	sessionRepo sessionRepo
	userGetter  userGetter
}

func NewSessionManager(
	sessionRepo sessionRepo,
	userGetter userGetter,
) *Manager {
	return &Manager{
		sessionRepo: sessionRepo,
		userGetter:  userGetter,
	}
}
