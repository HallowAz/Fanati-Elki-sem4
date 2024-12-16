package problem

import (
	"fe-sem4/config"
	"net/http"

	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
)

func (h *Handler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get(config.CookieHeader)
	if cookie == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	_, err := h.sessionStorer.GetSession(r.Context(), cookie)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	vars := mux.Vars(r)
	idStr := vars[idParam]

	id, err := tools.StrToUint32(idStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	err = h.problemStorer.DeleteProblem(r.Context(), id)
	if err != nil {
		processError(w, err)
	}

	return
}
