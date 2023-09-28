package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func reponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Reponding with 5xx error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	reponseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func reponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
