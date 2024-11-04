package session

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

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type sessionManager interface {
	Login(ctx context.Context, checkUser models.User) (string, error)
}

type SessionHandler struct {
	sessionManager sessionManager
}

func NewSessionHandler(sessionManager sessionManager) *SessionHandler {
	return &SessionHandler{
		sessionManager: sessionManager,
	}
}

func (h *SessionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/session", h.Login).Methods(http.MethodPost)
}

func processError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain_error.ErrCredentialsIncorrect):
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	log.Println(err)
	_ = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
	w.WriteHeader(http.StatusInternalServerError)

	return
}
