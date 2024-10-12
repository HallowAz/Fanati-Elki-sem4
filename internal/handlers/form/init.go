package form

// Для каждого публичного метода отдельный файл, поскольку так легче искать и теститься
import (
	"context"

	models "fe-sem4/internal/models/form"
)

type formManager interface {
	// Здесь все методы слоя менеджеров
	// Если слой менеджера для функции содержит чисто вызов следующего слоя, то не надо
	// для него делать слой менеджера, создай отдельный интерфейс для таких методов и вызывай
	// сразу слой репы
}

type problemStorer interface {
	CreateForm(ctx context.Context, problem models.Problem) error
	GetProblems(ctx context.Context) ([]*models.Problem, error)
}

type Handler struct {
	formManager   formManager
	problemStorer problemStorer
}

func NewFormHandler(formManager formManager, problemStorer problemStorer) *Handler {
	return &Handler{formManager: formManager, problemStorer: problemStorer}
}
