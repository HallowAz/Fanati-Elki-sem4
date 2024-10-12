package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "fe-sem4/internal/models/form"
	"fe-sem4/internal/repository/internal/form"
)

type ProblemRepo struct {
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{
		DB: db,
	}
}

// Это слой репы, здесь открывается транзакция и вызываются методы походов в базу и парсится ответ из бд
// Здесь не надо разделять методы на файлы, поскольку это вызовет потом проблемы

func (f *ProblemRepo) CreateForm(ctx context.Context, problem models.Problem) error {
	err := form.CreateForm(ctx, f.DB, problem)
	if err != nil {
		return fmt.Errorf("failed to create form: %w", err)
	}

	return nil
}

func (f *ProblemRepo) GetProblems(ctx context.Context) ([]*models.Problem, error) {
	problems, err := form.GetForms(ctx, f.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to create form: %w", err)
	}

	return problems, nil
}
