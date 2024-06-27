package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// Two really important helper functions that are used in every handler to give a appropriate response to the client

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server internal error detected")
	}
	type RespondError struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, RespondError{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, errJSON := json.Marshal(payload)
	if errJSON != nil {
		log.Printf("Error: %s", errJSON)
		return
	}
	w.WriteHeader(code)
	_, err := w.Write(jsonData)

	if err != nil {
		log.Printf("Could not write: %s", err)
		return
	}

}
