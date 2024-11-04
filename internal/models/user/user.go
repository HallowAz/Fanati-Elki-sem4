package user

import "time"

type User struct {
	ID        uint32
	Username  string
	Password  string
	Phone     string
	IconURL   string
	Icon      []byte
	BirthDate time.Time
	Gender    string
	IsAdmin   bool
}
