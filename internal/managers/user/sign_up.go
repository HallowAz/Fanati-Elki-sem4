package user

import (
	"context"
	"errors"
	"fmt"

	"fe-sem4/internal/models/domain_error"
	"fe-sem4/internal/models/user"
)

func (m *Manager) SignUp(ctx context.Context, newUser user.User) error {
	_, err := m.userStorer.GetUserByPhone(ctx, newUser.Phone)
	if err != nil && !errors.Is(err, domain_error.ErrUserNotFound) {
		return fmt.Errorf("failed to find user by phone: %w", err)
	}

	if !errors.Is(err, domain_error.ErrUserNotFound) {
		return domain_error.ErrUserAlreadyExist
	}

	return m.userStorer.CreateUser(ctx, newUser)
}
