package problem

import (
	"context"

	modelLib "fe-sem4/internal/models/problem"
	"github.com/google/uuid"
)

func (m *Manager) CreateProblem(ctx context.Context, model modelLib.Problem) error {
	model.Media = make([]string, 0, len(model.Media))

	for i := 0; i < len(model.MediaFiles); i++ {
		model.Media = append(model.Media, uuid.New().String())
	}

	return m.formStorer.CreateForm(ctx, model)
}
