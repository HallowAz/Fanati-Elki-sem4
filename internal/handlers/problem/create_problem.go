package problem

import (
	"encoding/json"
	"io"
	"net/http"
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

func (h *Handler) CreateForm(w http.ResponseWriter, r *http.Request) {
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

	problemDTO := CreateProblemRequest{}

	err = json.Unmarshal(body, &problemDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = json.NewEncoder(w).Encode(&Error{Err: "problems while unmarshalling JSON"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	err = h.problemStorer.CreateForm(r.Context(), problemDTO.ToModel())
	if err != nil {
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}
