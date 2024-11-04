package repository

import (
	"context"
	"errors"
	"fe-sem4/config"
	"fe-sem4/internal/models/domain_error"
	"fe-sem4/internal/models/session"
	"fe-sem4/internal/tools"
	"github.com/go-redis/redis"
)

const (
	keyPrefix = "session:"
)

type SessionRepo struct {
	cli *redis.Client
}

func NewSessionRepo(cli *redis.Client) *SessionRepo {
	return &SessionRepo{cli: cli}
}

func (s *SessionRepo) CreateSession(_ context.Context, sess session.Session) error {
	return s.cli.Set(keyPrefix+sess.Key, sess.UserID, config.SessionExpTime).Err()
}

func (s *SessionRepo) GetSession(_ context.Context, key string) (session.Session, error) {
	userIDStr, err := s.cli.Get(keyPrefix + key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return session.Session{}, domain_error.ErrSessionNotFound
		}

		return session.Session{}, err
	}

	userID, err := tools.StrToUint32(userIDStr)
	if err != nil {
		return session.Session{}, err
	}

	return session.Session{
		Key:    key,
		UserID: userID,
	}, nil
}
