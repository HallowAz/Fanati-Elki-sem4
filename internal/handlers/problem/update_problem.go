package problem

import (
	"encoding/json"
	"io"
	"net/http"

	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
)

func (h *Handler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	//cookie := r.Header.Get(config.CookieHeader)
	//if cookie == "" {
	//	w.WriteHeader(http.StatusUnauthorized)
	//
	//	return
	//}
	//
	//_, err := h.sessionStorer.GetSession(r.Context(), cookie)
	//if err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars[idParam]

	id, err := tools.StrToUint32(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = json.NewEncoder(w).Encode(&Error{Err: "problems while reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

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
		processError(w, err)
	}

	return
}
