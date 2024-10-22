package session

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	sessionmodels "fe-sem4/internal/models/session"
	usermodels "fe-sem4/internal/models/user"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ss Manager) Login(ctx context.Context, user *usermodels.User) (*sessionmodels.Cookie, error) {
	us, err := ss.userStorer.FindUserByPhone(ctx, user.Phone)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by phone: %w", err)
	}

	if us == nil {
		return nil, sessionmodels.ErrBadRequest
	}

	if user.Password != us.Password {
		return nil, sessionmodels.ErrBadRequest
	}

	cookie := &sessionmodels.Cookie{
		UserID:       us.ID,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge:       150 * time.Hour,
	}

	err = ss.sessionStorer.Create(cookie)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	return cookie, nil

}

func (ss Manager) Check(ctx context.Context, SessionToken string) (uint, error) {
	cookie, err := ss.sessionStorer.Check(SessionToken)
	if err != nil {
		return 0, err
	}
	if cookie == nil {
		return 0, nil
	}
	user, err := ss.userStorer.FindUserByID(ctx, cookie.UserID)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, nil
	}
	return user.ID, nil
}

func (ss Manager) Logout(ctx context.Context, cookie *sessionmodels.Cookie) error {
	return ss.sessionStorer.Delete(cookie)
}
