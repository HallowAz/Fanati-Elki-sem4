package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	models "fe-sem4/internal/models/problem"
	"fe-sem4/internal/repository/db"
	problem_db "fe-sem4/internal/repository/internal/problem"
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

func (f *ProblemRepo) CreateForm(ctx context.Context, problem models.Problem) error {
	var (
		problemRow = problem_db.NewProblemRow(problem)
		err        error
	)

	err = f.db.InTx(ctx, func(ctx context.Context, tx pgx.Tx) error {
		return problem_db.CreateForm(ctx, tx, problemRow)
	})
	if err != nil {
		return fmt.Errorf("failed to create problem: %w", err)
	}

	return nil
}

//func (f *ProblemRepo) GetProblems(ctx context.Context) ([]*models.Problem, error) {
//	problems, err := problem.GetForms(ctx, f.DB)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create form: %w", err)
//	}
//
//	return problems, nil
//}
