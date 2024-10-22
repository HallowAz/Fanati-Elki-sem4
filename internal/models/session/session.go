package session

import (
	"errors"
	"time"
)

var (
	ErrBadRequest = errors.New("bad request")
)

type Cookie struct {
	UserID       uint
	SessionToken string
	MaxAge       time.Duration
}
