package session

import (
	"encoding/json"
	"fe-sem4/config"
	"io"
	"net/http"
)

func (s *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems while reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	req := loginRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems while unmarshal;ing data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	cookie, err := s.sessionManager.Login(r.Context(), req.toModel())
	if err != nil {
		processError(w, err)

		return
	}

	w.Header().Set(config.CookieHeader, cookie)
}
