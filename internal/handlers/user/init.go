package user

import (
	"context"
	"encoding/json"
	"errors"
	"fe-sem4/internal/models/domain_error"
	models "fe-sem4/internal/models/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const idParam = "id"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type userManager interface {
	SignUp(ctx context.Context, newUser models.User) (string, error)
}

type userStorer interface {
	GetUserByID(ctx context.Context, id uint32) (models.User, error)
}

type Handler struct {
	userManager userManager
	userStorer  userStorer
}

func NewUserHandler(
	userManager userManager,
	userStorer userStorer,
) *Handler {
	return &Handler{
		userManager: userManager,
		userStorer:  userStorer,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods(http.MethodGet)
}

func processError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain_error.ErrUserAlreadyExist):
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return

	case errors.Is(err, domain_error.ErrUserNotFound):
		w.WriteHeader(http.StatusNotFound)
	}

	log.Println(err)
	_ = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)

	return
}
