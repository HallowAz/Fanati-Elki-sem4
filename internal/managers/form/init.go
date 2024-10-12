package form

import "context"

// Также все методы в отдельный файлах в том же пакете. Название файла - название метода

type formStorer interface {
	CreateForm(ctx context.Context)
}

type Manager struct {
	formStorer formStorer
}

func NewManager(formStorer formStorer) *Manager {
	return &Manager{
		formStorer: formStorer,
	}
}
