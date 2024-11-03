package user

import (
	"fe-sem4/internal/models/user"
	"time"
)

type createUserRequest struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Icon     []byte    `json:"icon"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	IsAdmin  bool      `json:"isAdmin"`
}

func (r *createUserRequest) ToModel() user.User {
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

type getUserByIDResponse struct {
	ID       uint32    `json:"id"`
	Username string    `json:"username"`
	Phone    string    `json:"phone"`
	Icon     []byte    `json:"icon"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	IsAdmin  bool      `json:"isAdmin"`
}

func newGetUserByIDResponse(model user.User) *getUserByIDResponse {
	return &getUserByIDResponse{
		ID:       model.ID,
		Username: model.Username,
		Phone:    model.Phone,
		Icon:     model.Icon,
		Birthday: model.BirthDate,
		Gender:   model.Gender,
		IsAdmin:  model.IsAdmin,
	}
}
