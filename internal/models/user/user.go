package user

import "errors"

var (
	ErrConflictPhoneNumber = errors.New("phone already exists")
)

type User struct {
	ID       uint
	Username string
	Password string
	Phone    string
	Icon     string
	Age      int
	Gender   string
	IsAdmin  bool
}
