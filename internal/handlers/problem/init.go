package problem

// Для каждого публичного метода отдельный файл, поскольку так легче искать и теститься
import (
	"context"
	"encoding/json"
	"errors"
	"fe-sem4/internal/models/domain_error"
	"net/http"

	models "fe-sem4/internal/models/problem"
)

type formManager interface {
	// Здесь все методы слоя менеджеров
	// Если слой менеджера для функции содержит чисто вызов следующего слоя, то не надо
	// для него делать слой менеджера, создай отдельный интерфейс для таких методов и вызывай
	// сразу слой репы
}

type problemStorer interface {
	CreateForm(ctx context.Context, problem models.Problem) error
	GetProblems(ctx context.Context) ([]models.Problem, error)
	UpdateProblem(ctx context.Context, problem models.Problem) error
}

type Handler struct {
	formManager   formManager
	problemStorer problemStorer
}

func NewFormHandler(formManager formManager, problemStorer problemStorer) *Handler {
	return &Handler{formManager: formManager, problemStorer: problemStorer}
}

func processError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain_error.ErrProblemNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)
}
