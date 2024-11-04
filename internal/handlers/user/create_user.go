package user

import (
	"encoding/json"
	"fe-sem4/config"
	"io"
	"net/http"
	"time"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	userDTO := createUserRequest{}
	err = json.Unmarshal(body, &userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems while unmarshalling json"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	cookie, err := h.userManager.SignUp(r.Context(), userDTO.ToModel())
	if err != nil {
		processError(w, err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     config.CookieName,
		Value:    cookie,
		Expires:  time.Now().Add(config.SessionExpTime),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}
