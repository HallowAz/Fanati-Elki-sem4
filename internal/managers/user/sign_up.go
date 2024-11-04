package user

import (
	"context"
	"errors"
	"fe-sem4/config"
	"fe-sem4/internal/models/domain_error"
	"fe-sem4/internal/models/session"
	"fe-sem4/internal/models/user"
	"fe-sem4/internal/tools"
)

func (m *Manager) SignUp(ctx context.Context, newUser user.User) (string, error) {
	_, err := m.userStorer.GetUserByPhone(ctx, newUser.Phone)
	if err != nil && !errors.Is(err, domain_error.ErrUserNotFound) {
		return "", nil
	}

	if !errors.Is(err, domain_error.ErrUserNotFound) {
		return "", domain_error.ErrUserAlreadyExist
	}

	err = m.userStorer.CreateUser(ctx, newUser)
	if err != nil {
		return "", err
	}

	cookie := tools.GenerateRandomString(config.CookieLen)

	return cookie, m.sessionCreator.CreateSession(ctx, session.Session{
		Key:    cookie,
		UserID: newUser.ID,
	})
}
