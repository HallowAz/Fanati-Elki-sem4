package user

import (
	"database/sql"
	"fe-sem4/internal/models/user"
	"time"
)

type UserRow struct {
	ID        uint32         `db:"id"`
	Username  string         `db:"username"`
	Password  string         `db:"password"`
	Phone     string         `db:"phone"`
	IconURL   sql.NullString `db:"icon_url"`
	BirthDate time.Time      `db:"birth_date"`
	Gender    sql.NullString `db:"gender"`
	IsAdmin   bool           `db:"is_admin"`
}

func NewUserRow(model user.User) UserRow {
	row := UserRow{
		ID:        model.ID,
		Username:  model.Username,
		Password:  model.Password,
		Phone:     model.Phone,
		BirthDate: model.BirthDate,
		IsAdmin:   model.IsAdmin,
	}

	if model.IconURL != "" {
		row.IconURL = sql.NullString{String: model.IconURL, Valid: true}
	}

	if model.Gender != "" {
		row.Gender = sql.NullString{String: model.Gender, Valid: true}
	}

	return row
}

func (r *UserRow) ToModel() user.User {
	return user.User{
		ID:        r.ID,
		Username:  r.Username,
		Password:  r.Password,
		Phone:     r.Phone,
		IconURL:   r.IconURL.String,
		BirthDate: r.BirthDate,
		Gender:    r.Gender.String,
		IsAdmin:   r.IsAdmin,
	}
}
