package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, status_code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("An error %v occured", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status_code)
	w.Write(data)
}
