package problem

//
//import (
//	"context"
//	"encoding/json"
//	"net/http"
//	"time"
//)
//
//func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
//	start := time.Now()
//	defer func() {
//		observeRequest(time.Since(start), "500")
//	}()
//
//	w.Header().Set("Content-Type", "application/json")
//
//	problems, err := h.problemStorer.GetProblems(context.Background())
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//	}
//
//	encoder := json.NewEncoder(w)
//	err = encoder.Encode(&Result{Body: problems})
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		err = json.NewEncoder(w).Encode(&Error{Err: "error while marshalling JSON"})
//		return
//	}
//
//}
