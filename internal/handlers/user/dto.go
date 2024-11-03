package user

import (
	"fe-sem4/internal/models/user"
	"time"
)

type CreateUserRequest struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Icon     []byte    `json:"icon"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	IsAdmin  bool      `json:"isAdmin"`
}

func (r CreateUserRequest) ToModel() user.User {
	return user.User{
		Username:  r.Username,
		Password:  r.Password,
		Phone:     r.Phone,
		Icon:      r.Icon,
		BirthDate: r.Birthday,
		Gender:    r.Gender,
		IsAdmin:   r.IsAdmin,
	}
}
