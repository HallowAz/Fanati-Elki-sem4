package problem

import (
	"bytes"
	"encoding/json"
	"fe-sem4/internal/tools"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
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

	//err = json.NewEncoder(w).Encode(&Result{Body: newGetProblemResponse(problem)})
	//if err != nil {
	//	processError(w, err)
	//
	//	return
	//}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Добавляем файл в multipart
	for i := 0; i < len(problem.Media); i++ {
		part, err := writer.CreateFormFile("files", problem.Media[i])
		if err != nil {
			log.Println(w, "Failed to create form file", http.StatusInternalServerError)
			return
		}

		_, err = io.Copy(part, bytes.NewReader(problem.MediaFiles[i]))
		if err != nil {
			log.Println(w, "Failed to copy file data", http.StatusInternalServerError)
			return
		}
	}
	// Добавляем JSON-ответ как текстовую часть
	resp := newGetProblemResponse(problem)

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	// Добавляем JSON как часть multipart
	partJSON, err := writer.CreateFormField("response")
	if err != nil {
		log.Println(w, "Failed to create form field for JSON", http.StatusInternalServerError)
		return
	}

	_, err = partJSON.Write(jsonResponse)
	if err != nil {
		log.Println(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}

	// Закрываем writer и устанавливаем заголовок Content-Type
	err = writer.Close()

	w.Header().Set("Content-Type", writer.FormDataContentType())

	// Отправляем ответ
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Println(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	return
}
