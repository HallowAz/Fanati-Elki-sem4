package repository

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gomodule/redigo/redis"

	models "fe-sem4/internal/models/session"
)

type SessionRepo struct {
	redisConn redis.Conn
	mu        sync.Mutex
}

func NewSessionRepo(conn redis.Conn) *SessionRepo {
	return &SessionRepo{
		redisConn: conn,
	}
}

func (sm *SessionRepo) Create(cookie *models.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "sessions:" + cookie.SessionToken
	sm.mu.Lock()
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	sm.mu.Unlock()
	if err != nil || result != "OK" {
		return fmt.Errorf("failed to set cookie: %w", err)
	}

	return nil
}

func (sm *SessionRepo) Check(sessionToken string) (*models.Cookie, error) {
	mkey := "sessions:" + sessionToken
	sm.mu.Lock()
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	sm.mu.Unlock()
	if err != nil {
		if err != redis.ErrNil {
			return nil, fmt.Errorf("failed to get cookie: %w", err)
		}
		return nil, nil
	}

	cookie := &models.Cookie{}
	err = json.Unmarshal(data, &cookie.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cookie: %w", err)
	}
	return cookie, nil
}

func (sm *SessionRepo) Delete(cookie *models.Cookie) error {
	mkey := "sessions:" + cookie.SessionToken
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		if err != redis.ErrNil {
			return fmt.Errorf("failed to delete cookie: %w", err)
		}
		return nil
	}

	return nil
}

func (sm *SessionRepo) Expire(cookie *models.Cookie) error {
	err := sm.Delete(cookie)
	if err != nil {
		return fmt.Errorf("failed to delete cookie: %w", err)
	}

	err = sm.Create(cookie)
	if err != nil {
		return fmt.Errorf("failed to create cookie: %w", err)
	}
	return nil
}
