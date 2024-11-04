package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, tx pgx.Tx, row UserRow) error {
	const query = `
			INSERT INTO users (
                   username,
                   password,
                   phone,
                   icon_url,
                   birth_date,
                   gender,
                   is_admin) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			  `

	_, err := tx.Exec(ctx, query,
		row.Username,
		row.Password,
		row.Phone,
		row.IconURL,
		row.BirthDate,
		row.Gender,
		row.IsAdmin)

	return err
}

func GetUserByID(ctx context.Context, tx pgx.Tx, id uint32) (UserRow, error) {
	const query = `
			SELECT id,
    			   username,
    			   password,
    			   phone,
    			   icon_url,
    			   birth_date,
    			   gender,
    			   is_admin
			FROM users
			WHERE id = $1
			`

	rows, err := tx.Query(ctx, query, id)
	if err != nil {
		return UserRow{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserRow])
}

func GetUserByPhone(ctx context.Context, tx pgx.Tx, phone string) (UserRow, error) {
	const query = `
			SELECT id,
    			   username,
    			   password,
    			   phone,
    			   icon_url,
    			   birth_date,
    			   gender,
    			   is_admin
			FROM users
			WHERE phone = $1
			`

	rows, err := tx.Query(ctx, query, phone)
	if err != nil {
		return UserRow{}, err
	}

	return pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserRow])
}
