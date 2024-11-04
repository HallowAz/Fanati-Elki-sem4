package session

import (
	"fe-sem4/config"
	"net/http"
)

func (s *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie := r.Header.Get(config.CookieHeader)
	if cookie == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	err := s.sessionStorer.DeleteSession(r.Context(), cookie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Del(config.CookieHeader)
}
