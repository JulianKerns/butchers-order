package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/JulianKerns/butchers-order/internal/database"
	"github.com/JulianKerns/butchers-order/models"
)

func (cfg *HandlerConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Decoding the emmail and username sent via POST request

	type Parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		log.Println("Error: Could not decode the Request Body")
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}
	now := time.Now().Format(time.DateTime)
	User, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		Username:  params.Username,
		Email:     params.Email,
	})
	if err != nil {
		log.Println("Could not create the user record")
		respondWithError(w, http.StatusInternalServerError, "Could not create the new User")
		return
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseUsertoUser(User))

}
