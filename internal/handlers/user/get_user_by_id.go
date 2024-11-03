package user

import (
	"encoding/json"
	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars[idParam]

	id, err := tools.StrToUint32(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	user, err := h.userStorer.GetUserByID(r.Context(), id)
	if err != nil {
		processError(w, err)

		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: newGetUserByIDResponse(user)})
	if err != nil {
		processError(w, err)

		return
	}
}
