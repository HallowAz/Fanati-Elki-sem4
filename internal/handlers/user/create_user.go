package user

import (
	"encoding/json"
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

	err = h.userManager.SignUp(r.Context(), userDTO.ToModel())
	if err != nil {
		processError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}
