package form

import (
	"fmt"
	"net/http"
)

func (h *Handler) CreateForm(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
}
