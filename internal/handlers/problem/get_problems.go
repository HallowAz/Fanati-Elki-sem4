package problem

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	models "fe-sem4/internal/models/problem"
)

func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		observeRequest(time.Since(start), "500")
	}()

	w.Header().Set("Content-Type", "application/json")

	problems, err := h.problemStorer.GetProblems(r.Context())
	if err != nil {
		log.Println(err)
		processError(w, err)
	}

	err = json.NewEncoder(w).Encode(&Result{Body: h.repackGetProblems(problems)})
	if err != nil {
		processError(w, err)

		return
	}
}

func (h *Handler) repackGetProblems(problems []models.Problem) GetProblemsResponse {
	var resp = GetProblemsResponse{
		Problems: make([]getProblemResponse, 0, len(problems)),
	}

	for _, problem := range problems {
		resp.Problems = append(resp.Problems, newGetProblemResponse(problem))
	}

	return resp
}
