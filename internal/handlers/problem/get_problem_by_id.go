package problem

import (
	"encoding/json"
	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handler) GetProblemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars[idParam]

	id, err := tools.StrToUint32(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	problem, err := h.problemStorer.GetProblemByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		processError(w, err)
	}

	err = json.NewEncoder(w).Encode(&Result{Body: newGetProblemResponse(problem)})
	if err != nil {
		processError(w, err)

		return
	}
}
