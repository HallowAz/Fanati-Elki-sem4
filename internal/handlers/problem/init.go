package problem

// Для каждого публичного метода отдельный файл, поскольку так легче искать и теститься
import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"fe-sem4/internal/models/domain_error"
	models "fe-sem4/internal/models/problem"
	"github.com/gorilla/mux"
)

const idParam = "id"

type formManager interface {
	CreateProblem(ctx context.Context, model models.Problem) error
}

type problemStorer interface {
	CreateForm(ctx context.Context, problem models.Problem) error
	GetProblems(ctx context.Context) ([]models.Problem, error)
	UpdateProblem(ctx context.Context, problem models.Problem) error
	DeleteProblem(ctx context.Context, id uint32) error
	GetProblemByID(ctx context.Context, id uint32) (models.Problem, error)
}

type Handler struct {
	formManager   formManager
	problemStorer problemStorer
}

func NewFormHandler(formManager formManager, problemStorer problemStorer) *Handler {
	return &Handler{formManager: formManager, problemStorer: problemStorer}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/problems", h.CreateForm).Methods(http.MethodPost)
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
