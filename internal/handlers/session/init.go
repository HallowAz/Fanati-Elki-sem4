package session

import (
	"context"

	sessionmodels "fe-sem4/internal/models/session"
	usermodels "fe-sem4/internal/models/user"
)

type sessionManager interface {
	Login(ctx context.Context, user *usermodels.User) (*sessionmodels.Cookie, error)
	Logout(ctx context.Context, cookie *sessionmodels.Cookie) error
	Check(ctx context.Context, SessionToken string) (uint, error)
}

type Handler struct {
	sessionManager sessionManager
}

func NewSessionHandler(sessionManager sessionManager) *Handler {
	return &Handler{sessionManager: sessionManager}
}
