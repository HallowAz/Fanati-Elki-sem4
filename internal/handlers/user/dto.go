package user

import (
	"fe-sem4/internal/models/user"
	"time"
)

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Icon     []byte `json:"icon"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (r *createUserRequest) ToModel() user.User {
	birthday, _ := time.Parse("2006-01-02", r.Birthday)

	return user.User{
		Username:  r.Username,
		Password:  r.Password,
		Phone:     r.Phone,
		Icon:      r.Icon,
		BirthDate: birthday,
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
