package problem

// Для каждого публичного метода отдельный файл, поскольку так легче искать и теститься
import (
	"context"
	"encoding/json"
	"errors"
	"fe-sem4/internal/models/session"
	"log"
	"net/http"

	"fe-sem4/internal/models/domain_error"
	models "fe-sem4/internal/models/problem"
	"github.com/gorilla/mux"
)

const idParam = "id"

type problemManager interface {
	CreateProblem(ctx context.Context, model models.Problem) error
}

type problemStorer interface {
	CreateProblem(ctx context.Context, problem models.Problem) error
	GetProblems(ctx context.Context) ([]models.Problem, error)
	UpdateProblem(ctx context.Context, problem models.Problem) error
	DeleteProblem(ctx context.Context, id uint32) error
	GetProblemByID(ctx context.Context, id uint32) (models.Problem, error)
}

type sessionStorer interface {
	GetSession(ctx context.Context, key string) (session.Session, error)
	DeleteSession(_ context.Context, key string) error
}

type Handler struct {
	problemManager problemManager
	problemStorer  problemStorer
	sessionStorer  sessionStorer
}

func NewProblemHandler(
	problemManager problemManager,
	problemStorer problemStorer,
	sessionStorer sessionStorer,
) *Handler {
	return &Handler{
		problemManager: problemManager,
		problemStorer:  problemStorer,
		sessionStorer:  sessionStorer,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/problems", h.CreateProblem).Methods(http.MethodPost)
	router.HandleFunc("/problems", h.GetProblems).Methods(http.MethodGet)
	router.HandleFunc("/problems/{id}", h.GetProblemByID).Methods(http.MethodGet)
	router.HandleFunc("/problems/{id}", h.UpdateProblem).Methods(http.MethodPatch)
	router.HandleFunc("/problems/{id}", h.DeleteProblem).Methods(http.MethodDelete)
}

func processError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain_error.ErrProblemNotFound):
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println(err)
	_ = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)
}
