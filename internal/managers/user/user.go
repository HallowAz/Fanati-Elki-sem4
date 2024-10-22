package user

import (
	"context"
	"fmt"

	usermodels "fe-sem4/internal/models/user"
)

func (m *Manager) SignUp(ctx context.Context, newUser *usermodels.User) error {
	user, err := m.userStorer.FindUserByPhone(ctx, newUser.Phone)
	if err != nil {
		return fmt.Errorf("failed to find user by phone: %w", err)
	}

	if user != nil {
		return usermodels.ErrConflictPhoneNumber
	}

	err = m.userStorer.CreateUser(ctx, *newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
