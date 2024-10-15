package form

import (
	"context"
	"database/sql"
	"fmt"

	models "fe-sem4/internal/models/form"
)

// Здесь описываются чисто sql запросы и походы в бд,
// я бы хотел пользоваться pgx сразу, но если не успеваешь, то используй любой
// Тут чисто функции, без структуры, так будет проще, когда надо будет комбинировать

func CreateForm(ctx context.Context, db *sql.DB, problem models.Problem) error {
	const insertProblem = `
		INSERT INTO problems (title, description, specific_location, category, vote_count, lat, long)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(
		insertProblem,
		problem.Title,
		problem.Description,
		problem.SpecificLocation,
		problem.Category,
		//problem.Media,
		problem.VoteCount,
		problem.Lat,
		problem.Long,
	)
	if err != nil {
		return fmt.Errorf("create form query failed: %w", err)
	}

	return nil
}

func GetForms(ctx context.Context, db *sql.DB) ([]*models.Problem, error) {
	const selectProblems = `
		SELECT id, title, description, specific_location, category, vote_count, lat, long
		FROM problems`
	var problems []*models.Problem
	rows, err := db.Query(selectProblems)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		problem := &models.Problem{}
		err = rows.Scan(
			&problem.ID,
			&problem.Title,
			&problem.Description,
			&problem.SpecificLocation,
			&problem.Category,
			&problem.VoteCount,
			&problem.Lat,
			&problem.Long,
		)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to read problems: %w", err)
	}
	return problems, nil
}
