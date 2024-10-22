package session

import (
	"context"
	"encoding/json"
	"net/http"
)

func (handler *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id, err := handler.sessionManager.Check(context.Background(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, cookie)

	err = json.NewEncoder(w).Encode(&Result{Body: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
