package repository

import (
	"context"
	"errors"
	"fe-sem4/internal/models/domain_error"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"

	models "fe-sem4/internal/models/problem"
	"fe-sem4/internal/repository/db"
	problem_db "fe-sem4/internal/repository/internal/problem"
	"github.com/jackc/pgx/v5"
)

type ProblemRepo struct {
	db db.TxCommitter
}

func NewProblemRepo(db db.TxCommitter) *ProblemRepo {
	return &ProblemRepo{
		db: db,
	}
}

// Это слой репы, здесь открывается транзакция и вызываются методы походов в базу и парсится ответ из бд
// Здесь не надо разделять методы на файлы, поскольку это вызовет потом проблемы

func (f *ProblemRepo) CreateProblem(ctx context.Context, problem models.Problem) error {
	var (
		problemRow = problem_db.NewProblemRow(problem)
		err        error
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		err = problem_db.CreateProblem(ctx, tx, problemRow)
		if err != nil {
			return err
		}

		return problem_db.SaveMediaToLocalStorage(ctx, problem.Media, problem.MediaFiles)
	})
	if err != nil {
		return fmt.Errorf("failed to create problem: %w", err)
	}

	return nil
}

func (f *ProblemRepo) GetProblems(ctx context.Context) ([]models.Problem, error) {
	var (
		problemRows []problem_db.ProblemRow
		err         error
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		problemRows, err = problem_db.GetProblems(ctx, tx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get problems: %w", err)
	}

	problems := make([]models.Problem, 0, len(problemRows))

	for _, row := range problemRows {
		problems = append(problems, row.ToModel())
	}

	return problems, nil
}

func (f *ProblemRepo) UpdateProblem(ctx context.Context, problem models.Problem) error {
	var (
		row      = problem_db.NewProblemRow(problem)
		err      error
		affected pgconn.CommandTag
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		affected, err = problem_db.UpdateProblem(ctx, tx, row)
		return err
	})
	if err != nil {
		return fmt.Errorf("failed to update problem: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return domain_error.ErrProblemNotFound
	}

	return nil
}

func (f *ProblemRepo) DeleteProblem(ctx context.Context, id uint32) error {
	var (
		affected pgconn.CommandTag
		err      error
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		affected, err = problem_db.DeleteProblem(ctx, tx, id)
		return err
	})
	if err != nil {
		return fmt.Errorf("failed to delete problem: %w", err)
	}

	if affected.RowsAffected() == 0 {
		return domain_error.ErrProblemNotFound
	}

	return nil
}

func (f *ProblemRepo) GetProblemByID(ctx context.Context, id uint32) (models.Problem, error) {
	var (
		row problem_db.ProblemRow
		err error
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		row, err = problem_db.GetProblemById(ctx, tx, id)
		return err

		//row.MediaFiles, err = problem_db.GetMediaFromLocalStorage(ctx, row.Media)
		//return err
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Problem{}, domain_error.ErrProblemNotFound
		}

		return models.Problem{}, fmt.Errorf("failed to get problem: %w", err)
	}

	return row.ToModel(), nil
}
