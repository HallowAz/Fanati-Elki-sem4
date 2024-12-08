package problem

import (
	"context"

	models "fe-sem4/internal/models/problem"
)

// Также все методы в отдельный файлах в том же пакете. Название файла - название метода

type problemStorer interface {
	CreateProblem(ctx context.Context, problem models.Problem) error
}

type Manager struct {
	problemStorer problemStorer
}

func NewManager(problemStorer problemStorer) *Manager {
	return &Manager{
		problemStorer: problemStorer,
	}
}
