package utils

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, status_code int, msg string) {
	payload := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	data, _ := json.Marshal(payload)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status_code)
	w.Write(data)
}
