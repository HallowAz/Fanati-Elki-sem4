package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	models "fe-sem4/internal/models/user"
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonbody, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user := models.User{}
	err = json.Unmarshal(jsonbody, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with unmarshaling json"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = h.userManager.SignUp(context.Background(), &user)
	if err != nil {
		if err == models.ErrConflictPhoneNumber {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusInternalServerError)
	}

}
