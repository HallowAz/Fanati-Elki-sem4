package session

import (
	"encoding/json"
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

	session, err := s.sessionStorer.GetSession(r.Context(), cookie)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: authResponse{UserID: session.UserID}})
	if err != nil {
		processError(w, err)

		return
	}

}
