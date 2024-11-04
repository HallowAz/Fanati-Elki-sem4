package session

import (
	"fe-sem4/config"
	"net/http"
)

func (s *SessionHandler) Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie := r.Header.Get(config.CookieHeader)
	if cookie == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	_, err := s.sessionGetter.GetSession(r.Context(), cookie)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
