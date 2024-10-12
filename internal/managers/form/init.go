package form

import (
	"context"

	models "fe-sem4/internal/models/form"
)

// Также все методы в отдельный файлах в том же пакете. Название файла - название метода

type formStorer interface {
	CreateForm(ctx context.Context, problem models.Problem) error
}

type Manager struct {
	formStorer formStorer
}

func NewManager(formStorer formStorer) *Manager {
	return &Manager{
		formStorer: formStorer,
	}
}
