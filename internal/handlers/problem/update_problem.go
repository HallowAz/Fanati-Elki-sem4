package problem

import (
	"encoding/json"
	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

const idParam = "id"

func (h *Handler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	idStr := vars[idParam]

	id, err := tools.StrToUint32(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	problemDTO := updateProblemRequest{
		ID: id,
	}

	err = json.Unmarshal(body, &problemDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = json.NewEncoder(w).Encode(&Error{Err: "problems while unmarshalling JSON"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	err = h.problemStorer.UpdateProblem(r.Context(), problemDTO.toModel())
	if err != nil {
		log.Println(err)

		processError(w, err)
	}

	return
}
