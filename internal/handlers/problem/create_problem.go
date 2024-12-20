package problem

import (
	"log"
	"net/http"
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

func (h *Handler) CreateProblem(w http.ResponseWriter, r *http.Request) {
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

	const maxFormSize = 16 << 20

	err := r.ParseMultipartForm(maxFormSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error parsing form:", err)
		return
	}

	problemDTO, err := parseFormCreateProblem(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error parsing form:", err)
		return
	}

	err = h.problemManager.CreateProblem(r.Context(), problemDTO.toModel())
	if err != nil {
		processError(w, err)
	}

	return
}
