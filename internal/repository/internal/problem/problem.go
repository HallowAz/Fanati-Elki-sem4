package problem

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Здесь описываются чисто sql запросы и походы в бд,
// я бы хотел пользоваться pgx сразу, но если не успеваешь, то используй любой
// Тут чисто функции, без структуры, так будет проще, когда надо будет комбинировать

const (
	dir    = "public/media/"
	extJPG = ".jpg"
)

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
		FROM problems
		WHERE is_deleted = false`

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
		WHERE id = $9 AND is_deleted = false`

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

func DeleteProblem(ctx context.Context, tx pgx.Tx, id uint32) (pgconn.CommandTag, error) {
	const query = `
		UPDATE problems
		SET is_deleted = true
		WHERE id = $1 AND is_deleted = false`

	affected, err := tx.Exec(ctx, query, id)

	return affected, err
}

func GetProblemById(ctx context.Context, tx pgx.Tx, id uint32) (ProblemRow, error) {
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
		FROM problems
		WHERE id = $1 AND is_deleted = false`

	row, err := tx.Query(ctx, query, id)
	if err != nil {
		return ProblemRow{}, err
	}

	return pgx.CollectOneRow(row, pgx.RowToStructByNameLax[ProblemRow])
}

func SaveMediaToLocalStorage(_ context.Context, filenames []string, images [][]byte) error {
	for i := 0; i < len(filenames); i++ {
		filename := dir + filenames[i] + extJPG

		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("error creating img file: %w", err)
		}

		_, err = file.Write(images[i])
		if err != nil {
			return fmt.Errorf("failed to save img file %s: %w", filename, err)
		}

		err = file.Close()
		if err != nil {
			log.Printf("failed to close img file %s: %v", filename, err)
		}
	}

	return nil
}

func GetMediaFromLocalStorage(_ context.Context, filenames []string) ([][]byte, error) {
	var images = make([][]byte, 0, len(filenames))

	for i := 0; i < len(filenames); i++ {
		filename := dir + filenames[i] + extJPG

		file, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("error opening img file %s: %w", filename, err)
		}

		image, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("error reading img file %s: %w", filename, err)
		}

		images = append(images, image)

		err = file.Close()
		if err != nil {
			log.Printf("failed to close img file %s: %v", filename, err)
		}
	}

	return images, nil
}
