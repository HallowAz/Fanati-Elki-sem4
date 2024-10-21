package problem

import (
	"context"
	"github.com/jackc/pgx/v5"
)

// Здесь описываются чисто sql запросы и походы в бд,
// я бы хотел пользоваться pgx сразу, но если не успеваешь, то используй любой
// Тут чисто функции, без структуры, так будет проще, когда надо будет комбинировать

func CreateForm(ctx context.Context, tx pgx.Tx, problemRow ProblemRow) error {
	const query = `
		INSERT INTO 
		    problems (
		              title,
		              description,
		              specific_location,
		              category,
		              media,
		              lat,
		              long
		              )
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := tx.Exec(
		ctx,
		query,
		problemRow.Title,
		problemRow.Description,
		problemRow.SpecificLocation,
		problemRow.Category,
		problemRow.Media,
		problemRow.Lat,
		problemRow.Long,
	)

	return err
}

//func GetForms(ctx context.Context, db *sql.DB) ([]*models.Problem, error) {
//	const selectProblems = `
//		SELECT id, title, description, specific_location, category, vote_count, lat, long
//		FROM problems`
//	var problems []*models.Problem
//	rows, err := db.Query(selectProblems)
//	if err != nil {
//		return nil, fmt.Errorf("query failed: %w", err)
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		problem := &models.Problem{}
//		err = rows.Scan(
//			&problem.ID,
//			&problem.Title,
//			&problem.Description,
//			&problem.SpecificLocation,
//			&problem.Category,
//			&problem.VoteCount,
//			&problem.Lat,
//			&problem.Long,
//		)
//		if err != nil {
//			return nil, err
//		}
//		problems = append(problems, problem)
//	}
//	err = rows.Err()
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return nil, nil
//		}
//		return nil, fmt.Errorf("failed to read problems: %w", err)
//	}
//	return problems, nil
//}
