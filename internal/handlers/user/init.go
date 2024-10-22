package user

import (
	"context"

	models "fe-sem4/internal/models/user"
)

type userManager interface {
	SignUp(ctx context.Context, user *models.User) error
}
type Handler struct {
	userManager userManager
}

func NewUserHandler(userManager userManager) *Handler {
	return &Handler{userManager: userManager}
}
