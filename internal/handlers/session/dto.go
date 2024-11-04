package session

import "fe-sem4/internal/models/user"

type loginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (l *loginRequest) toModel() user.User {
	return user.User{
		Phone:    l.Phone,
		Password: l.Password,
	}
}
