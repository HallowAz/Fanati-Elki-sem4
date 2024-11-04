package user

import (
	"encoding/json"
	"fe-sem4/config"
	"io"
	"net/http"
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

	req := createUserRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems while unmarshalling json"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	cookie, err := h.userManager.SignUp(r.Context(), req.ToModel())
	if err != nil {
		processError(w, err)
	}

	w.Header().Set(config.CookieHeader, cookie)
	w.WriteHeader(http.StatusCreated)
}
