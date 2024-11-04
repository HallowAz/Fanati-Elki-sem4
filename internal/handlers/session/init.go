package session

import (
	"context"
	"encoding/json"
	"errors"
	"fe-sem4/internal/models/domain_error"
	"fe-sem4/internal/models/session"
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

type sessionStorer interface {
	GetSession(ctx context.Context, key string) (session.Session, error)
	DeleteSession(_ context.Context, key string) error
}

type SessionHandler struct {
	sessionManager sessionManager
	sessionStorer  sessionStorer
}

func NewSessionHandler(
	sessionManager sessionManager,
	sessionStorer sessionStorer,
) *SessionHandler {
	return &SessionHandler{
		sessionManager: sessionManager,
		sessionStorer:  sessionStorer,
	}
}

func (s *SessionHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/session", s.Login).Methods(http.MethodPost)
	router.HandleFunc("/session", s.Auth).Methods(http.MethodGet)
	router.HandleFunc("/session", s.Logout).Methods(http.MethodDelete)
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
