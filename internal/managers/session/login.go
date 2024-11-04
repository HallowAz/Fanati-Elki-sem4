package session

import (
	"context"
	"fe-sem4/config"
	"fe-sem4/internal/models/domain_error"
	"fe-sem4/internal/models/session"
	models "fe-sem4/internal/models/user"
	"fe-sem4/internal/tools"
)

func (m *Manager) Login(ctx context.Context, checkUser models.User) (string, error) {
	user, err := m.userGetter.GetUserByPhone(ctx, checkUser.Phone)
	if err != nil {
		return "", err
	}

	if user.Password != checkUser.Password {
		return "", domain_error.ErrCredentialsIncorrect
	}

	cookie := tools.GenerateRandomString(config.CookieLen)

	err = m.sessionRepo.CreateSession(ctx, session.Session{
		Key:    cookie,
		UserID: user.ID,
	})
	if err != nil {
		return "", err
	}

	return cookie, nil
}
