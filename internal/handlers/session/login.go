package session

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	sessionmodels "fe-sem4/internal/models/session"
	usermodels "fe-sem4/internal/models/user"
)

const allowedOrigin = "http://84.23.53.216"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	jsonbody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	reqUser := usermodels.User{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with unmarshaling data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	cookieUC, err := handler.sessionManager.Login(context.Background(), &reqUser)
	if err != nil {
		if err == sessionmodels.ErrBadRequest {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    cookieUC.SessionToken,
		Expires:  time.Now().Add(cookieUC.MaxAge),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	err = json.NewEncoder(w).Encode(&Result{Body: reqUser.Phone})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
