package user

import (
	"context"
	"database/sql"
	"fmt"

	models "fe-sem4/internal/models/user"
)

func CreateUser(ctx context.Context, db *sql.DB, user *models.User) error {
	insertUser := `INSERT INTO users (username, password, phone, icon, age, gender, is_admin) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(insertUser, user.Username, user.Password, user.Phone, user.Icon, user.Age, user.Gender, user.IsAdmin)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
