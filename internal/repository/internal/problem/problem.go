package problem

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Здесь описываются чисто sql запросы и походы в бд,
// я бы хотел пользоваться pgx сразу, но если не успеваешь, то используй любой
// Тут чисто функции, без структуры, так будет проще, когда надо будет комбинировать

func CreateProblem(ctx context.Context, tx pgx.Tx, problemRow ProblemRow) error {
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

func GetProblems(ctx context.Context, tx pgx.Tx) ([]ProblemRow, error) {
	const query = `
		SELECT
		    id,
		    title,
		    description,
		    specific_location,
		    category,
		    media,
		    vote_count,
		    lat,
		    long
		FROM problems`

	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[ProblemRow])
}

func UpdateProblem(ctx context.Context, tx pgx.Tx, problemRow ProblemRow) (pgconn.CommandTag, error) {
	const query = `
		UPDATE problems
	    SET title = $1,
	        description = $2,
	        specific_location = $3,
	        category = $4,
	        media = $5,
	        vote_count = $6,
	        lat = $7,
	        long = $8
		WHERE id = $9`

	affected, err := tx.Exec(ctx, query,
		problemRow.Title,
		problemRow.Description,
		problemRow.SpecificLocation,
		problemRow.Category,
		problemRow.Media,
		problemRow.VoteCount,
		problemRow.Lat,
		problemRow.Long,
		problemRow.ID)

	return affected, err
}
