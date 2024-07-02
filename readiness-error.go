package main

import (
	"net/http"
)

// These handlers checks if the server is ready and serving endpoints.

// Tests if the respondWithJSON function is working as intened
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "This endpoint is reachable")
}

// Tests if the respondWithError function is working as intened
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "This endpoint is reachable but we respond with an Error")
}
