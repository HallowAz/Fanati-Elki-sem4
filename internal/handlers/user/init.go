package user

import (
	"context"

	models "fe-sem4/internal/models/user"
)

type userManager interface {
}

type userStorer interface {
	CreateUser(ctx context.Context, user models.User) error
}

type Handler struct {
	userManager userManager
	userStorer  userStorer
}

func NewUserHandler(userManager userManager, userStorer userStorer) *Handler {
	return &Handler{userManager: userManager, userStorer: userStorer}
}
