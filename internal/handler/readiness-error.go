package handler

import (
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "This endpoint is reachable")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "This endpoint is reachable but we respond with an Error")
}
